package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
)

func operAddUser(ctx context.Context, req auth.NewUserRequest) (*auth.User, error) {
	acc, ok := GetAccount(req.AccountID)
	if !ok {
		return nil, errors.Errorf("unknown account_id")
	}
	u, err := acc.AddUser(req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add user")
	}
	return &auth.User{
		AccountID: acc.id,
		ID:        u.id,
		Name:      u.name,
	}, nil
} //operAddUser()
