// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"sync"

	"v.io/v23/context"
	"v.io/v23/discovery"
)

type idiscovery struct {
	plugins []Plugin

	mu     sync.Mutex
	closed bool                  // GUARDED_BY(mu)
	tasks  map[*context.T]func() // GUARDED_BY(mu)
	wg     sync.WaitGroup

	adMu          sync.Mutex
	adSessions    map[discovery.AdId]sessionId  // GUARDED_BY(adMu)
	adSubtasks    map[discovery.AdId]*adSubtask // GUARDED_BY(adMu)
	adStopTrigger *Trigger

	dirServer *dirServer
}

type sessionId uint64

type adSubtask struct {
	parent *context.T

	mu   sync.Mutex
	stop func() // GUARDED_BY(mu)
}

func (d *idiscovery) shutdown() {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()
		return
	}
	d.dirServer.shutdown()
	for _, cancel := range d.tasks {
		cancel()
	}
	d.closed = true
	d.mu.Unlock()
	d.wg.Wait()

	for _, plugin := range d.plugins {
		plugin.Close()
	}
}

func (d *idiscovery) addTask(ctx *context.T) (*context.T, func(), error) {
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()
		return nil, nil, NewErrDiscoveryClosed(ctx)
	}
	ctx, cancel := context.WithCancel(ctx)
	d.tasks[ctx] = cancel
	d.wg.Add(1)
	d.mu.Unlock()
	return ctx, cancel, nil
}

func (d *idiscovery) removeTask(ctx *context.T) {
	d.mu.Lock()
	if _, exist := d.tasks[ctx]; exist {
		delete(d.tasks, ctx)
		d.wg.Done()
	}
	d.mu.Unlock()
}

func (d *idiscovery) cancelTask(ctx *context.T) {
	d.mu.Lock()
	cancel := d.tasks[ctx]
	d.mu.Unlock()
	if cancel != nil {
		cancel()
	}
}

func newDiscovery(ctx *context.T, plugins []Plugin) (*idiscovery, error) {
	if len(plugins) == 0 {
		return nil, NewErrNoDiscoveryPlugin(ctx)
	}
	d := &idiscovery{
		plugins:       make([]Plugin, len(plugins)),
		tasks:         make(map[*context.T]func()),
		adSessions:    make(map[discovery.AdId]sessionId),
		adSubtasks:    make(map[discovery.AdId]*adSubtask),
		adStopTrigger: NewTrigger(),
	}
	copy(d.plugins, plugins)

	// TODO(jhahn): Consider to start a directory server when it is required.
	// For example, scan-only applications would not need it.
	var err error
	if d.dirServer, err = newDirServer(ctx, d); err != nil {
		return nil, err
	}
	return d, nil
}
