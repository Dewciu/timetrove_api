package users

import (
	"github.com/dewciu/timetrove_api/pkg/common"
)

func GetAllUsersQuery() ([]UserModel, error) {
	var users []UserModel
	err := common.DB.Find(&users).Error
	return users, err
}

func CreateUser(user UserModel) error {
	return common.DB.Create(&user).Error
}
