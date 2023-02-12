package db

import "github.com/go-msvc/errors"

type UserRoleRecord struct {
	UserID string `db:"user_id"`
	RoleID string `db:"role_id"`
}

func AddUserRole(userID string, roleID string) error {
	u, err := GetUser(userID)
	if err != nil {
		return errors.Wrapf(err, "cannot get user")
	}
	r, err := GetRole(roleID)
	if err != nil {
		return errors.Wrapf(err, "cannot get role")
	}
	if u.AccountID == "" || r.AccountID == "" || u.AccountID != r.AccountID {
		return errors.Errorf("different accounts for user(%s).account(%s) and role(%s).account(%s)", u.ID, u.AccountID, r.ID, r.AccountID)
	}
	ur := UserRoleRecord{
		UserID: u.ID,
		RoleID: r.ID,
	}
	if _, err := db.Exec("INSERT INTO `user_roles` SET `user_id`=?,`role_id`=?", ur.UserID, ur.RoleID); err != nil {
		return errors.Wrapf(err, "failed to add user role") //e.g. duplicate name or broken foreign key
	}
	return nil
} //AddUserRole()
