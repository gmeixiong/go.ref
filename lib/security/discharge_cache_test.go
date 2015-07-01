// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"testing"
	"time"

	"v.io/v23/security"
	"v.io/v23/vdl"
)

func testDischargeCache(t *testing.T, s security.BlessingStore) {
	var (
		discharger = mkPrincipal()
		expiredCav = mkCaveat(security.NewPublicKeyCaveat(discharger.PublicKey(), "moline", security.ThirdPartyRequirements{}, security.UnconstrainedUse()))
		argsCav    = mkCaveat(security.NewPublicKeyCaveat(discharger.PublicKey(), "peoria", security.ThirdPartyRequirements{ReportArguments: true}, security.UnconstrainedUse()))
		methodCav  = mkCaveat(security.NewPublicKeyCaveat(discharger.PublicKey(), "moline", security.ThirdPartyRequirements{ReportMethod: true}, security.UnconstrainedUse()))
		serverCav  = mkCaveat(security.NewPublicKeyCaveat(discharger.PublicKey(), "peoria", security.ThirdPartyRequirements{ReportServer: true}, security.UnconstrainedUse()))

		dEmpty   = security.Discharge{}
		dExpired = mkDischarge(discharger.MintDischarge(expiredCav, mkCaveat(security.NewExpiryCaveat(time.Now().Add(-1*time.Minute)))))
		dArgs    = mkDischarge(discharger.MintDischarge(argsCav, security.UnconstrainedUse()))
		dMethod  = mkDischarge(discharger.MintDischarge(methodCav, security.UnconstrainedUse()))
		dServer  = mkDischarge(discharger.MintDischarge(serverCav, security.UnconstrainedUse()))

		emptyImp       = security.DischargeImpetus{}
		argsImp        = security.DischargeImpetus{Arguments: []*vdl.Value{&vdl.Value{}}}
		methodImp      = security.DischargeImpetus{Method: "foo"}
		otherMethodImp = security.DischargeImpetus{Method: "bar"}
		serverImp      = security.DischargeImpetus{Server: []security.BlessingPattern{security.BlessingPattern("fooserver")}}
		otherServerImp = security.DischargeImpetus{Server: []security.BlessingPattern{security.BlessingPattern("barserver")}}
	)

	// Discharges for different cavs should not be cached.
	d := mkDischarge(discharger.MintDischarge(argsCav, security.UnconstrainedUse()))
	s.CacheDischarge(d, argsCav, emptyImp)
	if d := s.Discharge(methodCav, emptyImp); d.ID() != "" {
		t.Errorf("Discharge for different caveat should not have been in cache")
	}
	s.ClearDischarges(d)

	// Add some discharges into the cache.
	s.CacheDischarge(dArgs, argsCav, argsImp)
	s.CacheDischarge(dMethod, methodCav, methodImp)
	s.CacheDischarge(dServer, serverCav, serverImp)
	s.CacheDischarge(dExpired, expiredCav, emptyImp)

	testCases := []struct {
		caveat          security.Caveat           // caveat that we are fetching discharges for.
		queryImpetus    security.DischargeImpetus // Impetus used to  query the cache.
		cachedDischarge security.Discharge        // Discharge that we expect to be returned from the cache, nil if the discharge should not be cached.
	}{
		// Expired discharges should not be returned by the cache.
		{expiredCav, emptyImp, dEmpty},

		// Discharges with Impetuses that have Arguments should not be cached.
		{argsCav, argsImp, dEmpty},

		{methodCav, methodImp, dMethod},
		{methodCav, otherMethodImp, dEmpty},
		{methodCav, emptyImp, dEmpty},

		{serverCav, serverImp, dServer},
		{serverCav, otherServerImp, dEmpty},
		{serverCav, emptyImp, dEmpty},
	}

	for i, test := range testCases {
		out := s.Discharge(test.caveat, test.queryImpetus)
		if got := out.ID(); got != test.cachedDischarge.ID() {
			t.Errorf("#%d: got discharge %v, want %v, queried with %v", i, got, test.cachedDischarge.ID(), test.queryImpetus)
		}
	}
	if t.Failed() {
		t.Logf("dArgs.ID():    %v", dArgs.ID())
		t.Logf("dMethod.ID():  %v", dMethod.ID())
		t.Logf("dServer.ID():  %v", dServer.ID())
		t.Logf("dExpired.ID(): %v", dExpired.ID())
	}
}

func mkPrincipal() security.Principal {
	p, err := NewPrincipal()
	if err != nil {
		panic(err)
	}
	return p
}

func mkDischarge(d security.Discharge, err error) security.Discharge {
	if err != nil {
		panic(err)
	}
	return d
}

func mkCaveat(c security.Caveat, err error) security.Caveat {
	if err != nil {
		panic(err)
	}
	return c
}