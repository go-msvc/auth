package auth

import "github.com/go-msvc/errors"

type Account struct {
	ID          string                `json:"id"`
	Users       []UserListEntry       `json:"users,omitempty"`
	Roles       []RoleListEntry       `json:"roles,omitempty"`
	Permissions []PermissionListEntry `json:"permissions,omitempty"`
}

type NewAccountRequest struct{}

type GetAccountRequest struct {
	ID       string `json:"id"`
	ListName string `json:"list_name" doc:"One of users|roles|permissions"`
	Filter   string `json:"list_filter,omitempty" doc:"Name filter is part of name, not a pattern"`
	Limit    int    `json:"list_limit" doc:"Limit 1..1000, default 10"`
}

func (req *GetAccountRequest) Validate() error {
	if req.ID == "" {
		return errors.Errorf("missing id")
	}
	switch req.ListName {
	case "users":
	case "roles":
	case "permissions":
	case "": //no list
	default:
		return errors.Errorf("unknown list_name:\"%s\" != users|roles|permissions", req.ListName)
	}
	if req.ListName == "" {
		if req.Limit != 0 {
			return errors.Errorf("specifying limit:%d without list_name", req.Limit)
		}
		if req.Filter != "" {
			return errors.Errorf("specifying filter:\"%s\" without list_name", req.Filter)
		}
	} else {
		if req.Limit == 0 {
			req.Limit = 10
		}
		if req.Limit < 1 || req.Limit > 1000 {
			return errors.Errorf("limit:%d not in valid range 1..1000", req.Limit)
		}
	}
	return nil
}

type DelAccountRequest struct {
	ID string `json:"id"`
}
