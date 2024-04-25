package database

import (
	"github.com/dewciu/timetrove_api/pkg/address"
	"github.com/dewciu/timetrove_api/pkg/users"
)

func Migrate() error {
	if err := DB.AutoMigrate(&users.User{}, &address.Address{}); err != nil {
		return err
	}

	return nil
}
