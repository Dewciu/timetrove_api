package database

import (
	"github.com/dewciu/timetrove_api/pkg/database/models"
)

func Migrate() error {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Address{},
	); err != nil {
		return err
	}

	return nil
}
