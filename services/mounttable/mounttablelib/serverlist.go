// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mounttablelib

import (
	"container/list"
	"sync"
	"time"

	"v.io/v23/naming"
	vdltime "v.io/v23/vdlroot/time"
)

type serverListClock interface {
	now() time.Time
}

type realTime bool

func (t realTime) now() time.Time {
	return time.Now()
}

// TODO(caprita): Replace this with the timekeeper library.
var slc = serverListClock(realTime(true))

// server maintains the state of a single server.  Unless expires is refreshed before the
// time is reached, the entry will be removed.
type server struct {
	expires time.Time
	oa      string // object address of server
}

// serverList represents an ordered list of servers.
type serverList struct {
	sync.Mutex
	l *list.List // contains entries of type *server
}

// newServerList creates a synchronized list of servers.
func newServerList() *serverList {
	return &serverList{l: list.New()}
}

// set up an alternate clock.
func setServerListClock(x serverListClock) {
	slc = x
}

func (sl *serverList) len() int {
	sl.Lock()
	defer sl.Unlock()
	return sl.l.Len()
}

func (sl *serverList) Front() *server {
	sl.Lock()
	defer sl.Unlock()
	return sl.l.Front().Value.(*server)
}

// add to the front of the list if not already in the list, otherwise,
// update the expiration time and move to the front of the list.  That
// way the most recently refreshed is always first.
func (sl *serverList) add(oa string, ttl time.Duration) {
	expires := slc.now().Add(ttl)
	sl.Lock()
	defer sl.Unlock()
	for e := sl.l.Front(); e != nil; e = e.Next() {
		s := e.Value.(*server)
		if s.oa == oa {
			s.expires = expires
			sl.l.MoveToFront(e)
			return
		}
	}
	s := &server{
		oa:      oa,
		expires: expires,
	}
	sl.l.PushFront(s) // innocent until proven guilty
}

// remove an element from the list.  Return the number of elements remaining.
func (sl *serverList) remove(oa string) int {
	sl.Lock()
	defer sl.Unlock()
	for e := sl.l.Front(); e != nil; e = e.Next() {
		s := e.Value.(*server)
		if s.oa == oa {
			sl.l.Remove(e)
			break
		}
	}
	return sl.l.Len()
}

// removeExpired removes any expired servers.
func (sl *serverList) removeExpired() int {
	sl.Lock()
	defer sl.Unlock()

	now := slc.now()
	var next *list.Element
	for e := sl.l.Front(); e != nil; e = next {
		s := e.Value.(*server)
		next = e.Next()
		if now.After(s.expires) {
			sl.l.Remove(e)
		}
	}
	return sl.l.Len()
}

// copyToSlice returns the contents of the list as a slice of MountedServer.
func (sl *serverList) copyToSlice() []naming.MountedServer {
	sl.Lock()
	defer sl.Unlock()
	var slice []naming.MountedServer
	for e := sl.l.Front(); e != nil; e = e.Next() {
		s := e.Value.(*server)
		ms := naming.MountedServer{
			Server:   s.oa,
			Deadline: vdltime.Deadline{s.expires},
		}
		slice = append(slice, ms)
	}
	return slice
}