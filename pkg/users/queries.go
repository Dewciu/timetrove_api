package users

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetAllUsersQuery() ([]UserModel, error) {
	var users []UserModel
	err := common.DB.Find(&users).Error
	return users, err
}
func CreateUserQuery(user UserModel) error {
	r := common.DB.Create(&user)

	fmt.Println(user)

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
