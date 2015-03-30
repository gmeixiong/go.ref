// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/verror"

	"v.io/x/lib/vlog"
)

var (
	errNoLocalBlessings = verror.Register("v.io/x/ref/services/security/roled/internal/noLocalBlessings", verror.NoRetry, "{1:}{2:} no local blessings")
)

type roleService struct {
	role               string
	config             *Config
	dischargerLocation string
}

func (i *roleService) SeekBlessings(call rpc.ServerCall) (security.Blessings, error) {
	ctx := call.Context()
	remoteBlessingNames, _ := security.RemoteBlessingNames(ctx)
	vlog.Infof("%q.SeekBlessings() called by %q", i.role, remoteBlessingNames)

	members := i.filterNonMembers(remoteBlessingNames)
	if len(members) == 0 {
		// The Authorizer should already have caught that.
		return security.Blessings{}, verror.New(verror.ErrNoAccess, ctx)
	}

	extensions := extensions(i.config, i.role, members)
	caveats, err := caveats(ctx, i.config)
	if err != nil {
		return security.Blessings{}, err
	}

	return createBlessings(ctx, i.config, v23.GetPrincipal(ctx), extensions, caveats, i.dischargerLocation)
}

// filterNonMembers returns only the blessing names that are authorized members
// for the role.
func (i *roleService) filterNonMembers(blessingNames []string) []string {
	var results []string
	for _, name := range blessingNames {
		// It is not enough to know if the pattern is matched by the
		// blessings. We need to know exactly which names matched.
		// These names will be used later to construct the role
		// blessings.
		for _, pattern := range i.config.Members {
			if pattern.MatchedBy(name) {
				results = append(results, name)
				break
			}
		}
	}
	return results
}

func extensions(config *Config, role string, blessingNames []string) []string {
	if !config.Extend {
		return []string{role}
	}
	var extensions []string
	for _, b := range blessingNames {
		extensions = append(extensions, role+security.ChainSeparator+b)
	}
	return extensions
}

func caveats(ctx *context.T, config *Config) ([]security.Caveat, error) {
	if config.Expiry == "" {
		return nil, nil
	}
	d, err := time.ParseDuration(config.Expiry)
	if err != nil {
		return nil, verror.Convert(verror.ErrInternal, ctx, err)
	}
	expiry, err := security.ExpiryCaveat(time.Now().Add(d))
	if err != nil {
		return nil, verror.Convert(verror.ErrInternal, ctx, err)
	}
	return []security.Caveat{expiry}, nil
}

func createBlessings(ctx *context.T, config *Config, principal security.Principal, extensions []string, caveats []security.Caveat, dischargerLocation string) (security.Blessings, error) {
	blessWith := security.GetCall(ctx).LocalBlessings()
	blessWithNames := security.LocalBlessingNames(ctx)
	publicKey := security.GetCall(ctx).RemoteBlessings().PublicKey()
	if len(blessWithNames) == 0 {
		return security.Blessings{}, verror.New(errNoLocalBlessings, ctx)
	}

	var ret security.Blessings
	for _, ext := range extensions {
		cav := caveats
		if config.Audit {
			// TODO(rthellend): This third-party caveat will only work with a single
			// discharger service. We need a way to allow multiple instances of this
			// service to be interchangeable.

			fullNames := make([]string, len(blessWithNames))
			for i, n := range blessWithNames {
				fullNames[i] = n + security.ChainSeparator + ext
			}
			loggingCaveat, err := security.NewCaveat(LoggingCaveat, fullNames)
			if err != nil {
				return security.Blessings{}, verror.Convert(verror.ErrInternal, ctx, err)
			}
			thirdParty, err := security.NewPublicKeyCaveat(principal.PublicKey(), dischargerLocation, security.ThirdPartyRequirements{true, true, true}, loggingCaveat)
			if err != nil {
				return security.Blessings{}, verror.Convert(verror.ErrInternal, ctx, err)
			}
			cav = append(cav, thirdParty)
		}
		if len(cav) == 0 {
			// TODO(rthellend,ashankar): the use of unconstrained
			// use is concerning. We should figure out how to get
			// rid of it.
			// Some options:
			//  - have the seeker specify a set of caveats in the
			//    request (and forcefully insert a restrictive one
			//    or fail if the role server thinks that they are
			//    too loose or something).
			//  - have a set of caveats in the config of the role.
			cav = []security.Caveat{security.UnconstrainedUse()}
		}
		b, err := principal.Bless(publicKey, blessWith, ext, cav[0], cav[1:]...)
		if err != nil {
			return security.Blessings{}, verror.Convert(verror.ErrInternal, ctx, err)
		}
		if ret, err = security.UnionOfBlessings(ret, b); err != nil {
			verror.Convert(verror.ErrInternal, ctx, err)
		}
	}
	return ret, nil
}
