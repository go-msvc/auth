package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/auth/db"
	"github.com/go-msvc/errors"
)

func operAddAccountUser(ctx context.Context, req auth.NewUserRequest) (*auth.User, error) {
	acc, err := db.GetAccount(req.AccountID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get account")
	}
	u, err := db.AddAccountUser(acc, req.Name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to add user")
	}
	return u, nil
} //operAddAccountUser()

func operAddUserRole(ctx context.Context, req auth.AddUserRoleRequest) error {
	if err := db.AddUserRole(req.UserID, req.RoleID); err != nil {
		return errors.Wrapf(err, "failed to add user role")
	}
	return nil
}
