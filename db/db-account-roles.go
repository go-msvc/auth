package db

import (
	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
	"github.com/google/uuid"
)

type AccountRoleRecord struct {
	AccountID string `db:"account_id"`
	ID        string `db:"id"`
	Name      string `db:"name"`
}

func AddAccountRole(acc *auth.Account, name string) (*auth.Role, error) {
	r := AccountRoleRecord{
		AccountID: acc.ID,
		ID:        uuid.New().String(),
		Name:      name,
	}
	if _, err := db.Exec("INSERT INTO `roles` SET `account_id`=?,`id`=?,`name`=?", r.AccountID, r.ID, r.Name); err != nil {
		return nil, errors.Wrapf(err, "failed to insert role") //e.g. duplicate name
	}
	return &auth.Role{
		AccountID: r.AccountID,
		ID:        r.ID,
		Name:      r.Name,
	}, nil
} //AddAccountRole()

func GetRole(id string) (*auth.Role, error) {
	var r AccountRoleRecord
	if err := db.Get(&r, "SELECT `account_id`,`id`,`name` FROM `roles` WHERE `id`=?", id); err != nil {
		return nil, errors.Wrapf(err, "failed to get role") //e.g. not found
	}
	return &auth.Role{
		AccountID: r.AccountID,
		ID:        r.ID,
		Name:      r.Name,
	}, nil
} //GetRole()

func DelRole(id string) error {
	if result, err := db.Exec("DELETE FROM `roles` WHERE `id`=?", id); err != nil {
		return errors.Wrapf(err, "failed to delete role") //e.g. not found
	} else if nr(result.RowsAffected()) != 1 {
		return errors.Wrapf(err, "did not delete role") //e.g. not found
	}
	return nil
} //DelRole()
