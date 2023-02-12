package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/auth/db"
	"github.com/go-msvc/errors"
)

func operAddAccountPermission(ctx context.Context, req auth.NewPermissionRequest) (*auth.Permission, error) {
	acc, err := db.GetAccount(req.AccountID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get account")
	}
	p, err := db.AddAccountPermission(acc, req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add permission")
	}
	return p, nil
} //operAddAccountPermission()
