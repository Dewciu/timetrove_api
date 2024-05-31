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
			fmt.Println(err.Detail)
			column := common.GetColumnFromUniqueErrorDetails(err.Detail)
			return &common.AlreadyExistsError{Column: column}
		}

		return err
	}

	return nil
}

func GetUsersFromContextQuery(c *gin.Context) ([]UserModel, error) {
	var users []UserModel
	query := common.DB

	if name := c.Query("name"); name != "" {
		query = query.Where("name = ?", name)
	}
	if email := c.Query("email"); email != "" {
		query = query.Where("email = ?", email)
	}
	if age := c.Query("age"); age != "" {
		query = query.Where("age = ?", age)
	}

	if err := query.Find(&users).Error; err != nil {
		return []UserModel{}, err
	}

	return users, nil
}
