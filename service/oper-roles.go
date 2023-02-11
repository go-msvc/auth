package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
)

func operAddRole(ctx context.Context, req auth.NewRoleRequest) (*auth.Role, error) {
	acc, ok := GetAccount(req.AccountID)
	if !ok {
		return nil, errors.Errorf("unknown account_id")
	}
	u, err := acc.AddRole(req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add role")
	}
	return &auth.Role{
		AccountID: acc.id,
		ID:        u.id,
		Name:      u.name,
	}, nil
} //operAddRole()
