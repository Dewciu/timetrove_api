package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username  string    `gorm:"unique;not null; type:varchar(100)" json:"username"`
	Email     string    `gorm:"unique;not null; type:varchar(100)" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	AddressID uuid.UUID `gorm:"default:null" json:"address_id"`
	Address   Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
} //@name User

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

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
