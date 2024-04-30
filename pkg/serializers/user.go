package serializers

import (
	"github.com/dewciu/timetrove_api/pkg/database/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID       `json:"id"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Address  AddressResponse `json:"address"`
} //@name User

type UserSerializer struct {
	C *gin.Context
	models.User
}

func (s *UserSerializer) Response() UserResponse {
	addressSerializer := AddressSerializer{s.C, s.Address}
	response := UserResponse{
		ID:       s.ID,
		Username: s.Username,
		Email:    s.Email,
		Address:  addressSerializer.Response(),
	}

	return response
}

type UsersSerializer struct {
	C     *gin.Context
	Users []models.User
}

func (s *UsersSerializer) Response() []UserResponse {
	var response []UserResponse
	for _, user := range s.Users {
		serializer := UserSerializer{s.C, user}
		response = append(response, serializer.Response())
	}

	return response
}
