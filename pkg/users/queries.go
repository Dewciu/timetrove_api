package users

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetAllUsersQuery() ([]UserModel, error) {
	var users []UserModel
	err := common.DB.Find(&users).Error
	return users, err
}
func CreateUserQuery(user UserModel) error {
	r := common.DB.Create(&user)

	if r.Error != nil {
		err := r.Error.(*pgconn.PgError)

		if err.Code == "23505" {
			column := common.GetColumnFromUniqueErrorDetails(err.Detail)
			return &common.AlreadyExistsError{Column: column}
		}

		return err
	}

	return nil
}

func GetUsersByFilterQuery(c *gin.Context) ([]UserModel, error) {
	var users []UserModel
	query := common.DB

	if username := c.Query("username"); username != "" {
		query = query.Where("username = ?", username)
	}
	if email := c.Query("email"); email != "" {
		query = query.Where("email = ?", email)
	}
	if id := c.Query("id"); id != "" {
		query = query.Where("id = ?", id)
	}

	if err := query.Find(&users).Error; err != nil {
		return []UserModel{}, err
	}

	return users, nil
}

func GetUserByIdQuery(id string) (UserModel, error) {
	var user UserModel
	err := common.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}
