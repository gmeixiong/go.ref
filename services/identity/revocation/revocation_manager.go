// Package revocation provides tools to create and manage revocation caveats.
package revocation

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"v.io/core/veyron2/security"
	"v.io/core/veyron2/vdl"
	"v.io/core/veyron2/vom"
)

// RevocationManager persists information for revocation caveats to provided discharges and allow for future revocations.
type RevocationManager interface {
	NewCaveat(discharger security.PublicKey, dischargerLocation string) (security.Caveat, error)
	Revoke(caveatID string) error
	GetRevocationTime(caveatID string) *time.Time
}

// revocationManager persists information for revocation caveats to provided discharges and allow for future revocations.
type revocationManager struct{}

// NewRevocationManager returns a RevocationManager that persists information about
// revocationCaveats in a SQL database and allows for revocation and caveat creation.
// This function can only be called once because of the use of global variables.
func NewRevocationManager(sqlDB *sql.DB) (RevocationManager, error) {
	revocationLock.Lock()
	defer revocationLock.Unlock()
	if revocationDB != nil {
		return nil, fmt.Errorf("NewRevocationManager can only be called once")
	}
	var err error
	revocationDB, err = newSQLDatabase(sqlDB, "RevocationCaveatInfo")
	if err != nil {
		return nil, err
	}
	return &revocationManager{}, nil
}

var revocationDB database
var revocationLock sync.RWMutex

// NewCaveat returns a security.Caveat constructed with a ThirdPartyCaveat for which discharges will be
// issued iff Revoke has not been called for the returned caveat.
func (r *revocationManager) NewCaveat(discharger security.PublicKey, dischargerLocation string) (security.Caveat, error) {
	var empty security.Caveat
	var revocation [16]byte
	if _, err := rand.Read(revocation[:]); err != nil {
		return empty, err
	}
	restriction, err := security.NewCaveat(NotRevokedCaveat, revocation[:])
	if err != nil {
		return empty, err
	}
	// TODO(ashankar): Remove when removing ValidatorVOM
	if restriction.ValidatorVOM, err = vom.Encode(revocationCaveat(revocation)); err != nil {
		return empty, err
	}
	cav, err := security.NewPublicKeyCaveat(discharger, dischargerLocation, security.ThirdPartyRequirements{}, restriction)
	if err != nil {
		return empty, err
	}
	if err = revocationDB.InsertCaveat(cav.ThirdPartyDetails().ID(), revocation[:]); err != nil {
		return empty, err
	}
	return cav, nil
}

// Revoke disables discharges from being issued for the provided third-party caveat.
func (r *revocationManager) Revoke(caveatID string) error {
	return revocationDB.Revoke(caveatID)
}

// GetRevocationTimestamp returns the timestamp at which a caveat was revoked.
// If the caveat wasn't revoked returns nil
func (r *revocationManager) GetRevocationTime(caveatID string) *time.Time {
	timestamp, err := revocationDB.RevocationTime(caveatID)
	if err != nil {
		return nil
	}
	return timestamp
}

func isRevoked(_ security.Context, key []byte) error {
	revocationLock.RLock()
	if revocationDB == nil {
		revocationLock.RUnlock()
		return fmt.Errorf("missing call to NewRevocationManager")
	}
	revocationLock.RUnlock()
	revoked, err := revocationDB.IsRevoked(key)
	if revoked {
		return fmt.Errorf("revoked")
	}
	return err
}

// TODO(ashankar): Remove when removing security.Caveat.ValidatorVOM.
type revocationCaveat [16]byte

func (cav revocationCaveat) Validate(ctx security.Context) error {
	return isRevoked(ctx, cav[:])
}

func init() {
	// TODO(ashankar): Remove when removing security.Caveat.ValidatorVOM
	vdl.Register(revocationCaveat{})
	security.RegisterCaveatValidator(NotRevokedCaveat, isRevoked)
}
