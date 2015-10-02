// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mdns

import (
	"encoding/base32"
	"errors"
	"regexp"
	"sort"
	"strings"
)

const (
	// Limit the maximum large txt records to the maximum total txt records size.
	maxLargeTxtRecordLen = maxTotalTxtRecordsLen
)

var (
	// The key of encoded large txt records is "_x<i><j>", where 'i' and 'j' will
	// be one digit numbers since we limit the large txt record to 1300 bytes.
	reLargeTxtRecord = regexp.MustCompile("^" + attrLargeTxtPrefix + "[0-9][0-9]=")

	errInvalidLargeTxtRecord = errors.New("invalid large txt record")
)

// encodeInstanceUuid encodes the given instance uuid to a valid host name by using
// "Extended Hex Alphabet" defined in RFC 4648. This removes any padding characters.
func encodeInstanceUuid(instanceUuid []byte) string {
	return strings.TrimRight(base32.HexEncoding.EncodeToString(instanceUuid), "=")
}

// decodeInstanceUuid decodes the given host name to an instance uuid.
func decodeInstanceUuid(hostname string) ([]byte, error) {
	// Add padding characters if needed.
	if p := len(hostname) % 8; p > 0 {
		hostname += strings.Repeat("=", 8-p)
	}
	return base32.HexEncoding.DecodeString(hostname)
}

// maybeSplitLargeTXT slices txt records larger than 255 bytes into multiple txt records.
func maybeSplitLargeTXT(txt []string) ([]string, error) {
	splitted := make([]string, 0, len(txt))
	xno := 0
	for _, v := range txt {
		switch n := len(v); {
		case n > maxLargeTxtRecordLen:
			return nil, errMaxTxtRecordLenExceeded
		case n > maxTxtRecordLen:
			var buf [maxTxtRecordLen]byte
			copy(buf[:], attrLargeTxtPrefix)
			for i, off := 0, 0; off < n; i++ {
				buf[2] = byte(xno + '0')
				buf[3] = byte(i + '0')
				buf[4] = '='
				c := copy(buf[5:], v[off:])
				splitted = append(splitted, string(buf[:5+c]))
				off += c
			}
			xno++
		default:
			splitted = append(splitted, v)
		}
	}
	return splitted, nil
}

// maybeJoinLargeTXT joins the splitted large txt records.
func maybeJoinLargeTXT(txt []string) ([]string, error) {
	joined, splitted := make([]string, 0, len(txt)), make([]string, 0)
	for _, v := range txt {
		switch {
		case strings.HasPrefix(v, attrLargeTxtPrefix):
			if !reLargeTxtRecord.MatchString(v) {
				return nil, errInvalidLargeTxtRecord
			}
			splitted = append(splitted, v)
		default:
			joined = append(joined, v)
		}
	}
	if len(splitted) == 0 {
		return joined, nil
	}

	sort.Strings(splitted)

	var buf [maxLargeTxtRecordLen]byte
	xno, off := 0, 0
	for _, v := range splitted {
		i := int(v[2] - '0')
		if i > xno {
			// A new large txt record started.
			joined = append(joined, string(buf[:off]))
			xno++
			off = 0
		}
		c := copy(buf[off:], v[5:])
		off += c
	}
	joined = append(joined, string(buf[:off]))
	return joined, nil
}
