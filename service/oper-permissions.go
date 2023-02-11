package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
)

func operAddPermission(ctx context.Context, req auth.NewPermissionRequest) (*auth.Permission, error) {
	acc, ok := GetAccount(req.AccountID)
	if !ok {
		return nil, errors.Errorf("unknown account_id")
	}
	p, err := acc.AddPermission(req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add permission")
	}
	return &auth.Permission{
		AccountID: acc.id,
		ID:        p.id,
		Name:      p.name,
	}, nil
} //operAddPermission()
