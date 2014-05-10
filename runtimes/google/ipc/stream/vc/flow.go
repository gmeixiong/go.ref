package vc

import (
	"net"
	"time"

	"veyron2/security"
)

type flow struct {
	idHolder
	*reader
	*writer
	laddr, raddr net.Addr
}

type idHolder interface {
	LocalID() security.PublicID
	RemoteID() security.PublicID
}

func (f *flow) LocalAddr() net.Addr  { return f.laddr }
func (f *flow) RemoteAddr() net.Addr { return f.raddr }
func (f *flow) Close() error {
	f.reader.Close()
	f.writer.Close()
	return nil
}

// SetDeadline sets a deadline on the flow. The flow will be cancelled if it
// is not closed by the specified deadline.
// A zero deadline (time.Time.IsZero) implies that no cancellation is desired.
func (f *flow) SetDeadline(t time.Time) error {
	if err := f.SetReadDeadline(t); err != nil {
		return err
	}
	if err := f.SetWriteDeadline(t); err != nil {
		return err
	}
	return nil
}

// Shutdown closes the flow and discards any queued up write buffers.
// This is appropriate when the flow has been closed by the remote end.
func (f *flow) Shutdown() {
	f.reader.Close()
	f.writer.Shutdown(true)
}

// Cancel closes the flow and discards any queued up write buffers.
// This is appropriate when the flow is being cancelled locally.
func (f *flow) Cancel() {
	f.reader.Close()
	f.writer.Shutdown(false)
}
