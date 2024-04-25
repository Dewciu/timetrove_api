package users

import (
	"github.com/dewciu/timetrove_api/pkg/address"
	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	common.BaseModel
	Username  string          `gorm:"unique;not null; type:varchar(100)" json:"username"`
	Email     string          `gorm:"unique;not null; type:varchar(100)" json:"email"`
	Password  string          `gorm:"not null" json:"password"`
	AddressID uuid.UUID       `gorm:"default:null"`
	Address   address.Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

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
