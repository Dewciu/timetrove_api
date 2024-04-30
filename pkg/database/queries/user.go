package queries

import (
	"github.com/dewciu/timetrove_api/pkg/database"
	"github.com/dewciu/timetrove_api/pkg/database/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}
