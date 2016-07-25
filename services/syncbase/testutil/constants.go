// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testutil

import (
	"strings"
)

var invalidBlessingPatterns = []string{
	",",
	"a,b",
	":",
	"a:",
	":b",
	"a::b",
	"a/b",
	"a/",
	"$",
}

var validBlessingPatterns = []string{
	"a",
	"a:b",
	"a:%",
	"a:%A",
	"v.io",
	"v.io:foo",
	"v.io:foobar",
	"v.io:foo:bar",
	"v.io:a:admin@myapp.com",
	"v.io:o:app:user",
	"\xfa",
	"a\xfb",
	"안녕하세요",
}

var validNames = []string{
	"a",
	"B",
	"a_",
	"a__",
	"a0_",
	"a_b",
	"a_0",
	"foo",
	"foobar",
	"foo_bar",
	"BARBAZ",
	"/",
	"a/b",
	":",
	"a:b",
	"*",
	"\xfa",
	"a\xfb",
	"@@",
	"dev.v.io/a/admin@myapp.com",
	"dev.v.io:a:admin@myapp.com",
	",",
	"a,b",
	" ",
	"foo bar",
	"foo-bar/baz42",
	"8T:j=re{*(TfF_.U9VG\"{2g:;Z-/~Ibm}&.p%7Zf:I{K~b8;Di\\!rA@k",
	"안녕하세요",
}

var longNames = []string{
	// 64 bytes
	"abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcd",
	// 16 4-byte runes => 64 bytes
	strings.Repeat("𠜎", 16),
	// 65 bytes
	"abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcde",
	// 16 4-byte runes + 1 more byte => 65 bytes
	strings.Repeat("𠜎", 16) + "a",
	// 256 4-byte runes => 1024 bytes
	strings.Repeat("𠜎", 256),
}

var veryLongNames = []string{
	// 256 4-byte runes + 1 more byte => 1025 bytes
	strings.Repeat("𠜎", 256) + "a",
	strings.Repeat("foobar", 1337),
}

var universallyInvalidIdParts = []string{
	// disallow empty
	"",
	// disallow Unicode control characters (00-1F, 7F, 80-9F)
	"\x00",
	"a\x00b",
	"\x01",
	"\x07",
	"\x1f",
	"\x7f",
	"\u0080",
	"\u0091",
	"\u009f",
	"a\x01b",
	"a\x7fb",
	"a\u0091b",
	// disallow bytes FC-FF (invalid Unicode)
	"\xfc",
	"\xfd",
	"\xfe",
	"\xff",
	"a\xfcb",
	"a\xfdb",
	"a\xfeb",
	"a\xffb",
}

func concat(slices ...[]string) []string {
	var res []string
	for _, slice := range slices {
		res = append(res, slice...)
	}
	return res
}

var (
	OkAppUserBlessings    []string = concat(validBlessingPatterns, longNames)
	NotOkAppUserBlessings []string = concat(universallyInvalidIdParts, invalidBlessingPatterns, veryLongNames)

	OkDbCxNames    []string = concat(validNames, longNames)
	NotOkDbCxNames []string = concat(universallyInvalidIdParts, veryLongNames)
)

var (
	OkRowKeys    []string = concat(validNames, longNames, veryLongNames)
	NotOkRowKeys []string = universallyInvalidIdParts
)
