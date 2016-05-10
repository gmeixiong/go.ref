// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/security"
	"v.io/v23/security/access"
	wire "v.io/v23/services/syncbase"
	"v.io/v23/syncbase"
	"v.io/v23/verror"
	"v.io/x/ref/services/syncbase/longevity_tests/model"
	"v.io/x/ref/services/syncbase/testutil"
)

// CreateDbsAndCollections creates databases and collections according to the
// given models.  It does not fail if any of the databases or collections
// already exist.  If the model contains syncgroups, it will also create or
// join those as well.
func CreateDbsAndCollections(ctx *context.T, sbName string, dbModels model.DatabaseSet) (map[syncbase.Database][]syncbase.Collection, []syncbase.Syncgroup, error) {
	nsRoots := v23.GetNamespace(ctx).Roots()

	service := syncbase.NewService(sbName)
	syncgroups := []syncbase.Syncgroup{}
	dbColsMap := map[syncbase.Database][]syncbase.Collection{}
	for _, dbModel := range dbModels {
		dbPerms := defaultPerms(ctx)
		if dbModel.Permissions != nil {
			dbPerms = dbModel.Permissions.ToWire("root")
		}
		allowChecker(dbPerms)

		// Create Database.
		// TODO(nlacasse): Don't create the database unless its blessings match
		// ours.
		db := service.DatabaseForId(dbModel.Id(), nil)
		if err := db.Create(ctx, dbPerms); err != nil && verror.ErrorID(err) != verror.ErrExist.ID {
			return nil, nil, err
		}
		dbColsMap[db] = []syncbase.Collection{}

		// Create collections for database.
		for _, colModel := range dbModel.Collections {
			colPerms := defaultPerms(ctx)
			if colModel.Permissions != nil {
				colPerms = colModel.Permissions.ToWire("root")
			}
			allowChecker(colPerms)

			// TODO(nlacasse): Don't create the collection unless its blessings
			// match ours.
			col := db.CollectionForId(colModel.Id())
			if err := col.Create(ctx, colPerms); err != nil && verror.ErrorID(err) != verror.ErrExist.ID {
				return nil, nil, err
			}
			dbColsMap[db] = append(dbColsMap[db], col)
		}

		// Create or join syncgroups for database.
		for _, sgModel := range dbModel.Syncgroups {
			sg := db.SyncgroupForId(wire.Id{Name: sgModel.NameSuffix, Blessing: "blessing"})
			if sgModel.HostDevice.Name == sbName {
				// We are the host.  Create the syncgroup.
				spec := sgModel.Spec("root")
				spec.MountTables = nsRoots

				if spec.Perms == nil {
					spec.Perms = defaultPerms(ctx)
				}
				allowChecker(spec.Perms)

				if err := sg.Create(ctx, spec, wire.SyncgroupMemberInfo{}); err != nil && verror.ErrorID(err) != verror.ErrExist.ID {
					return nil, nil, err
				}
				syncgroups = append(syncgroups, sg)
				continue
			}
			// Join the syncgroup.  It might not exist at first, so we loop.
			// TODO(nlacasse): Parameterize number of retries.  Exponential
			// backoff?
			var joinErr error
			for i := 0; i < 10; i++ {
				_, joinErr = sg.Join(ctx, sgModel.HostDevice.Name, "", wire.SyncgroupMemberInfo{})
				if joinErr == nil {
					syncgroups = append(syncgroups, sg)
					break
				} else {
					time.Sleep(100 * time.Millisecond)
				}
			}
			if joinErr != nil {
				return nil, nil, fmt.Errorf("could not join syncgroup %q: %v", sgModel.Name(), joinErr)
			}
		}
	}

	return dbColsMap, syncgroups, nil
}

// allowChecker gives the checker access to all tags on the Permissions object.
func allowChecker(perms access.Permissions) {
	checkerPattern := security.BlessingPattern("root:checker")
	for _, tag := range access.AllTypicalTags() {
		perms.Add(checkerPattern, string(tag))
	}
}

// defaultPerms returns a Permissions object that allows the context's default
// blessings.
func defaultPerms(ctx *context.T) access.Permissions {
	blessing, _ := v23.GetPrincipal(ctx).BlessingStore().Default()
	return testutil.DefaultPerms(blessing.String())
}
