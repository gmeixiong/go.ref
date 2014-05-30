package exec

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

var (
	ErrNoVersion          = errors.New(versionVariable + " environment variable missing")
	ErrUnsupportedVersion = errors.New("Unsupported version of veyron/lib/exec request by " + versionVariable + " environment variable")
)

type ChildHandle struct {
	// Endpoint is a callback endpoint that can be use to notify the
	// parent that the child has started up successfully via the
	// Callback() RPC.
	Endpoint string
	// ID is a callback ID that can be used by a parent to identify this
	// child when the child invokes the Callback() RPC using the
	// callback endpoint.
	ID string
	// Secret is a secret passed to the child by its parent via a
	// trusted channel.
	Secret string
	// statusPipe is a pipe that is used to notify the parent that the
	// child process has started successfully. Unlike the Callback()
	// RPC, which is to be invoked by the application to notify the
	// parent that the application is "ready", the statusPipe is to be
	// invoked by the veyron framework to notify the parent that the
	// child process has successfully started.
	statusPipe *os.File
}

// fileOffset accounts for the file descriptors that are always passed
// to the child by the parent: stderr, stdin, stdout, data read, and
// status write. Any extra files added by the client will follow
// fileOffset.
const fileOffset = 5

// NewChildHandle creates a new ChildHandle that can be used to signal
// that the child is 'ready' (by calling SetReady) to its parent. The
// value of the ChildHandle's Secret securely passed to it by the
// parent; this is intended for subsequent use to create a secure
// communication channels and or authentication.
//
// If the child is relying on exec.Cmd.ExtraFiles then its first file
// descriptor will not be 3, but will be offset by extra files added
// by the framework. The developer should use the NewExtraFile method
// to robustly get their extra files with the correct offset applied.
func NewChildHandle() (*ChildHandle, error) {
	switch os.Getenv(versionVariable) {
	case "":
		return nil, ErrNoVersion
	case version1:
		// TODO(cnicolaou): need to use major.minor.build format for
		// version #s.
	default:
		return nil, ErrUnsupportedVersion
	}
	dataPipe := os.NewFile(3, "data_rd")
	endpoint, err := readData(dataPipe)
	if err != nil {
		return nil, err
	}
	id, err := readData(dataPipe)
	if err != nil {
		return nil, err
	}
	secret, err := readData(dataPipe)
	if err != nil {
		return nil, err
	}
	c := &ChildHandle{
		Endpoint:   endpoint,
		ID:         id,
		Secret:     secret,
		statusPipe: os.NewFile(4, "status_wr"),
	}
	return c, nil
}

// SetReady writes a 'ready' status to its parent.
func (c *ChildHandle) SetReady() error {
	_, err := c.statusPipe.Write([]byte(readyStatus))
	c.statusPipe.Close()
	return err
}

// NewExtraFile creates a new file handle for the i-th file descriptor after
// discounting stdout, stderr, stdin and the files reserved by the framework for
// its own purposes.
func (c *ChildHandle) NewExtraFile(i uintptr, name string) *os.File {
	return os.NewFile(i+fileOffset, name)
}

func readData(r io.Reader) (string, error) {
	var l int64 = 0
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return "", err
	}
	var data []byte = make([]byte, l)
	if n, err := r.Read(data); err != nil || int64(n) != l {
		if err != nil {
			return "", err
		} else {
			return "", errors.New("partial read")
		}
	}
	return string(data), nil
}
