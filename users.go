package auth

import "github.com/go-msvc/errors"

type User struct {
	AccountID string          `json:"account_id"`
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Roles     []RoleListEntry `json:"roles,omitempty"`
}

type UserListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewUserRequest struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name" doc:"Name of new user"`
}

func (req NewUserRequest) Validate() error {
	if req.AccountID == "" {
		return errors.Errorf("missing account_id")
	}
	if req.Name == "" {
		return errors.Errorf("missing name")
	}
	return nil
}

type AddUserRoleRequest struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

func (req AddUserRoleRequest) Validate() error {
	if req.UserID == "" {
		return errors.Errorf("missing user_id")
	}
	if req.RoleID == "" {
		return errors.Errorf("missing role_id")
	}
	return nil
}
