package users

import (
	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserModelValidator struct {
	User struct {
		Username  string    `form:"username" json:"username" binding:"required,alphanum,min=4,max=255"`
		Email     string    `form:"email" json:"email" binding:"required,email"`
		Password  string    `form:"password" json:"password" binding:"required,min=8,max=255"`
		AddressID uuid.UUID `form:"address_id" json:"address_id"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (s *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}

	s.userModel.ID = uuid.New()
	s.userModel.Username = s.User.Username
	s.userModel.Email = s.User.Email
	s.userModel.Password = s.User.Password

	return nil
}
