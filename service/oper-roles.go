package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/auth/db"
	"github.com/go-msvc/errors"
)

func operAddAccountRole(ctx context.Context, req auth.NewRoleRequest) (*auth.Role, error) {
	acc, err := db.GetAccount(req.AccountID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get account")
	}
	r, err := db.AddAccountRole(acc, req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add role")
	}
	return r, nil
} //operAddAccountRole()

func operAddRolePermission(ctx context.Context, req auth.AddRolePermissionRequest) error {
	if err := db.AddRolePermission(req.RoleID, req.PermissionID); err != nil {
		return errors.Wrapf(err, "failed to add role permission")
	}
	return nil
} //operAddRolePermission()
