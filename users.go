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
