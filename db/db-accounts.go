package db

import (
	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
	"github.com/google/uuid"
)

type AccountRecord struct {
	ID     string `db:"id"`
	Active bool   `db:"active"`
	Name   string `db:"name"`
}

func AddAccount(name string) (*auth.Account, error) {
	acc := AccountRecord{
		ID:     uuid.New().String(),
		Name:   name,
		Active: true,
	}
	if _, err := db.Exec("INSERT INTO `accounts` SET `id`=?,`name`=?,`active`=?", acc.ID, acc.Name, acc.Active); err != nil {
		return nil, errors.Wrapf(err, "failed to insert account") //e.g. duplicate name
	}
	return &auth.Account{
		ID:     acc.ID,
		Name:   acc.Name,
		Active: acc.Active,
	}, nil
} //AddAccount()

func GetAccount(id string) (*auth.Account, error) {
	var acc AccountRecord
	if err := db.Get(&acc, "SELECT `id`,`name`,`active` FROM `accounts` WHERE `id`=?", id); err != nil {
		return nil, errors.Wrapf(err, "failed to get account") //e.g. not found
	}
	return &auth.Account{
		ID:     acc.ID,
		Name:   acc.Name,
		Active: acc.Active,
	}, nil
} //GetAccount()

func DelAccount(id string) error {
	if result, err := db.Exec("DELETE FROM `accounts` WHERE `id`=?", id); err != nil {
		return errors.Wrapf(err, "failed to delete account") //e.g. not found
	} else if nr(result.RowsAffected()) != 1 {
		return errors.Wrapf(err, "did not delete account") //e.g. not found
	}
	return nil
} //DelAccount()

func nr(nr int64, err error) int {
	if err != nil {
		return 0
	}
	return int(nr)
}
