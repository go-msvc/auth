package db

import (
	"github.com/go-msvc/auth"
	"github.com/go-msvc/errors"
	"github.com/google/uuid"
)

type AccountUserRecord struct {
	AccountID string `db:"account_id"`
	ID        string `db:"id"`
	Name      string `db:"name"`
}

func AddAccountUser(acc *auth.Account, name string) (*auth.User, error) {
	u := AccountUserRecord{
		AccountID: acc.ID,
		ID:        uuid.New().String(),
		Name:      name,
	}
	if _, err := db.Exec("INSERT INTO `users` SET `account_id`=?,`id`=?,`name`=?", u.AccountID, u.ID, u.Name); err != nil {
		return nil, errors.Wrapf(err, "failed to insert user") //e.g. duplicate name
	}
	return &auth.User{
		AccountID: u.AccountID,
		ID:        u.ID,
		Name:      u.Name,
	}, nil
} //AddAccountUser()

func GetUser(id string) (*auth.User, error) {
	var u AccountUserRecord
	if err := db.Get(&u, "SELECT `account_id`,`id`,`name` FROM `users` WHERE `id`=?", id); err != nil {
		return nil, errors.Wrapf(err, "failed to get user") //e.g. not found
	}
	return &auth.User{
		AccountID: u.AccountID,
		ID:        u.ID,
		Name:      u.Name,
	}, nil
} //GetUser()

func DelUser(id string) error {
	if result, err := db.Exec("DELETE FROM `users` WHERE `id`=?", id); err != nil {
		return errors.Wrapf(err, "failed to delete user") //e.g. not found
	} else if nr(result.RowsAffected()) != 1 {
		return errors.Wrapf(err, "did not delete user") //e.g. not found
	}
	return nil
} //DelUser()
