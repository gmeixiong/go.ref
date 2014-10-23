package blesser

import (
	"bytes"
	"fmt"
	"time"

	"veyron.io/veyron/veyron/services/identity"
	"veyron.io/veyron/veyron/services/identity/util"

	"veyron.io/veyron/veyron2"
	"veyron.io/veyron/veyron2/ipc"
	"veyron.io/veyron/veyron2/security"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
	"veyron.io/veyron/veyron2/vom"
)

type macaroonBlesser struct {
	rt  veyron2.Runtime // TODO(ashankar): Remove when the old security model is ripped out
	key []byte
}

// BlessingMacaroon contains the data that is encoded into the macaroon for creating blessings.
type BlessingMacaroon struct {
	Creation time.Time
	Caveats  []security.Caveat
	Name     string
}

// NewMacaroonBlesserServer provides an identity.MacaroonBlesser Service that generates blessings
// after unpacking a BlessingMacaroon.
//
// TODO(ashankar): Remove the "r" argument once the switch to the new security model is complete.
func NewMacaroonBlesserServer(r veyron2.Runtime, key []byte) interface{} {
	return identity.NewServerMacaroonBlesser(&macaroonBlesser{r, key})
}

func (b *macaroonBlesser) Bless(ctx ipc.ServerContext, macaroon string) (vdlutil.Any, error) {
	inputs, err := util.Macaroon(macaroon).Decode(b.key)
	if err != nil {
		return nil, err
	}
	var m BlessingMacaroon
	if err := vom.NewDecoder(bytes.NewBuffer(inputs)).Decode(&m); err != nil {
		return nil, err
	}
	if time.Now().After(m.Creation.Add(time.Minute * 5)) {
		return nil, fmt.Errorf("macaroon has expired")
	}
	if ctx.LocalPrincipal() == nil || ctx.RemoteBlessings() == nil {
		// TODO(ashankar): Old security model, remove this block.
		self := b.rt.Identity()
		var err error
		// Use the blessing that was used to authenticate with the client to bless it.
		if self, err = self.Derive(ctx.LocalID()); err != nil {
			return nil, err
		}
		return self.Bless(ctx.RemoteID(), m.Name, time.Hour*24*365, m.Caveats)
	}
	if len(m.Caveats) == 0 {
		m.Caveats = []security.Caveat{security.UnconstrainedUse()}
	}
	// TODO(ashankar,toddw): After the old security model is ripped out and the VDL configuration
	// files have the scheme to translate between "wire" types and "in-memory" types, this should just
	// become return ctx.LocalPrincipal().....
	blessings, err := ctx.LocalPrincipal().Bless(ctx.RemoteBlessings().PublicKey(), ctx.LocalBlessings(), m.Name, m.Caveats[0], m.Caveats[1:]...)
	if err != nil {
		return nil, err
	}
	return security.MarshalBlessings(blessings), nil
}
