// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command rpsbot repeatedly runs automated games, implementing all three roles.
// It publishes itself as player, judge, and scorekeeper. Then, it initiates
// games with other players, in a loop. As soon as one game is over, it starts a
// new one.
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/x/lib/vlog"
	"v.io/x/ref/examples/rps"
	"v.io/x/ref/examples/rps/internal"
	"v.io/x/ref/lib/security/securityflag"
	"v.io/x/ref/lib/signals"

	_ "v.io/x/ref/profiles/roaming"
)

var (
	name     = flag.String("name", "", "identifier to publish itself as (defaults to user@hostname)")
	numGames = flag.Int("num-games", -1, "number of games to play (-1 means unlimited)")
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()

	auth := securityflag.NewAuthorizerOrDie()
	server, err := v23.NewServer(ctx)
	if err != nil {
		vlog.Fatalf("NewServer failed: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	rpsService := NewRPS(ctx)

	listenSpec := v23.GetListenSpec(ctx)
	eps, err := server.Listen(listenSpec)
	if err != nil {
		vlog.Fatalf("Listen(%v) failed: %v", listenSpec, err)
	}
	if *name == "" {
		*name = internal.CreateName()
	}
	names := []string{
		fmt.Sprintf("rps/judge/%s", *name),
		fmt.Sprintf("rps/player/%s", *name),
		fmt.Sprintf("rps/scorekeeper/%s", *name),
	}
	if err := server.Serve(names[0], rps.RockPaperScissorsServer(rpsService), auth); err != nil {
		vlog.Fatalf("Serve(%v) failed: %v", names[0], err)
	}
	for _, n := range names[1:] {
		if err := server.AddName(n); err != nil {
			vlog.Fatalf("(%v) failed: %v", n, err)
		}
	}
	vlog.Infof("Listening on endpoint %s (published as %v)", eps, names)

	go initiateGames(ctx, rpsService)
	<-signals.ShutdownOnSignals(ctx)
}

func initiateGames(ctx *context.T, rpsService *RPS) {
	for i := 0; i < *numGames || *numGames == -1; i++ {
		if err := rpsService.Player().InitiateGame(ctx); err != nil {
			vlog.Infof("Failed to initiate game: %v", err)
		}
		time.Sleep(5 * time.Second)
	}
}
