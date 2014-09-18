package vc

import (
	"errors"

	"veyron.io/veyron/veyron/runtimes/google/lib/upcqueue"
	"veyron.io/veyron/veyron2/ipc/stream"
)

var errListenerClosed = errors.New("Listener has been closed")

type listener struct {
	q *upcqueue.T
}

func newListener() *listener { return &listener{q: upcqueue.New()} }

func (l *listener) Enqueue(f stream.Flow) error {
	err := l.q.Put(f)
	if err == upcqueue.ErrQueueIsClosed {
		return errListenerClosed
	}
	return err
}

func (l *listener) Accept() (stream.Flow, error) {
	item, err := l.q.Get(nil)
	if err == upcqueue.ErrQueueIsClosed {
		return nil, errListenerClosed
	}
	if err != nil {
		return nil, err
	}
	return item.(stream.Flow), nil
}

func (l *listener) Close() error {
	l.q.Close()
	return nil
}
