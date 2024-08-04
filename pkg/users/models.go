package users

import (
	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/dewciu/timetrove_api/pkg/middleware"
	"github.com/dewciu/timetrove_api/pkg/permissions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//TODO Add permissions to endpoints for the user

type UserModel struct {
	common.BaseModel
	Username    string                         `gorm:"unique;not null; type:varchar(255)" json:"username"`
	Email       string                         `gorm:"unique;not null; type:varchar(255)" json:"email"`
	Password    string                         `gorm:"not null" json:"password"`
	Token       string                         `gorm:"-" json:"token"`
	Permissions []*permissions.PermissionModel `gorm:"many2many:user_permissions;"`
} //@name User

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}

	hash, err := generatePassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hash

	return u.BaseModel.BeforeCreate(tx)
}

func (u *UserModel) LoginCheck() (string, error) {

	var user UserModel

	result := common.DB.Model(&UserModel{}).Where("username = ?", u.Username).First(&user)
	err := result.Error

	if err != nil {
		return "", err
	}

	err = verifyPassword(u.Password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := middleware.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
