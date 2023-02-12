package db

import "github.com/go-msvc/errors"

type RolePermissionRecord struct {
	RoleID       string `db:"role_id"`
	PermissionID string `db:"permission_id"`
}

func AddRolePermission(roleID string, permissionID string) error {
	r, err := GetRole(roleID)
	if err != nil {
		return errors.Wrapf(err, "cannot get role")
	}
	p, err := GetPermission(permissionID)
	if err != nil {
		return errors.Wrapf(err, "cannot get permission")
	}
	if r.AccountID == "" || p.AccountID == "" || r.AccountID != p.AccountID {
		return errors.Errorf("different accounts for role(%s).account(%s) and permission(%s).account(%s)", r.ID, r.AccountID, p.ID, p.AccountID)
	}
	rp := RolePermissionRecord{
		RoleID:       r.ID,
		PermissionID: p.ID,
	}
	if _, err := db.Exec("INSERT INTO `role_permissions` SET `role_id`=?,`permission_id`=?", rp.RoleID, rp.PermissionID); err != nil {
		return errors.Wrapf(err, "failed to add role permission") //e.g. duplicate name or broken foreign key
	}
	return nil
} //AddRolePermission()
