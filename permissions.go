package auth

import "github.com/go-msvc/errors"

type Permission struct {
	AccountID string `json:"account_id"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type PermissionListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewPermissionRequest struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name" doc:"Name of new permission"`
}

func (req NewPermissionRequest) Validate() error {
	if req.AccountID == "" {
		return errors.Errorf("missing account_id")
	}
	if req.Name == "" {
		return errors.Errorf("missing name")
	}
	return nil
}
