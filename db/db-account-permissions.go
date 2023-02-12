package db

import (
	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
	"github.com/google/uuid"
)

type AccountPermissionRecord struct {
	AccountID string `db:"account_id"`
	ID        string `db:"id"`
	Name      string `db:"name"`
}

func AddAccountPermission(acc *auth.Account, name string) (*auth.Permission, error) {
	p := AccountPermissionRecord{
		AccountID: acc.ID,
		ID:        uuid.New().String(),
		Name:      name,
	}
	if _, err := db.Exec("INSERT INTO `permissions` SET `account_id`=?,`id`=?,`name`=?", p.AccountID, p.ID, p.Name); err != nil {
		return nil, errors.Wrapf(err, "failed to insert permission") //e.g. duplicate name
	}
	return &auth.Permission{
		AccountID: p.AccountID,
		ID:        p.ID,
		Name:      p.Name,
	}, nil
} //AddAccountPermission()

func GetPermission(id string) (*auth.Permission, error) {
	var p AccountPermissionRecord
	if err := db.Get(&p, "SELECT `account_id`,`id`,`name` FROM `permissions` WHERE `id`=?", id); err != nil {
		return nil, errors.Wrapf(err, "failed to get permission") //e.g. not found
	}
	return &auth.Permission{
		AccountID: p.AccountID,
		ID:        p.ID,
		Name:      p.Name,
	}, nil
} //GetPermission()

func DelPermission(id string) error {
	if result, err := db.Exec("DELETE FROM `permissions` WHERE `id`=?", id); err != nil {
		return errors.Wrapf(err, "failed to delete permissions") //e.g. not found
	} else if nr(result.RowsAffected()) != 1 {
		return errors.Wrapf(err, "did not delete permissions") //e.g. not found
	}
	return nil
} //DelPermission()
