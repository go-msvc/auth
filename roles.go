package auth

import "github.com/go-msvc/errors"

type Role struct {
	AccountID   string                `json:"account_id"`
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Permissions []PermissionListEntry `json:"permissions,omitempty"`
}

type RoleListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewRoleRequest struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name" doc:"Name of new role"`
}

func (req NewRoleRequest) Validate() error {
	if req.AccountID == "" {
		return errors.Errorf("missing account_id")
	}
	if req.Name == "" {
		return errors.Errorf("missing name")
	}
	return nil
}

type AddRolePermissionRequest struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

func (req AddRolePermissionRequest) Validate() error {
	if req.RoleID == "" {
		return errors.Errorf("missing role_id")
	}
	if req.PermissionID == "" {
		return errors.Errorf("missing permission_id")
	}
	return nil
}
