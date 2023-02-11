package main

import (
	"context"

	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
)

func operAddAccount(ctx context.Context, req auth.NewAccountRequest) (res *auth.Account, err error) {
	acc, err := AddAccount()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create account")
	}
	return &auth.Account{
		ID: acc.id,
	}, nil
} //operAddAccount()

func operGetAccount(ctx context.Context, req auth.GetAccountRequest) (res *auth.Account, err error) {
	acc, ok := GetAccount(req.ID)
	if !ok {
		return nil, errors.Wrapf(err, "failed to get account")
	}
	res = &auth.Account{
		ID: acc.id,
	}
	switch req.ListName {
	case "users":
		users := acc.findUsers(req.Filter, req.Limit)
		for _, u := range users {
			res.Users = append(res.Users, auth.UserListEntry{Name: u.name, ID: u.id})
		}
	case "roles":
		roles := acc.findRoles(req.Filter, req.Limit)
		for _, r := range roles {
			res.Roles = append(res.Roles, auth.RoleListEntry{Name: r.name, ID: r.id})
		}
	case "permissions":
		permissions := acc.findPermissions(req.Filter, req.Limit)
		for _, p := range permissions {
			res.Permissions = append(res.Permissions, auth.PermissionListEntry{Name: p.name, ID: p.id})
		}
	default:
	}
	return res, nil
} //operGetAccount()

func operDelAccount(ctx context.Context, req auth.DelAccountRequest) (err error) {
	if err := DelAccount(req.ID); err != nil {
		return errors.Wrapf(err, "failed to delete account")
	}
	return nil
} //operDelAccount()
