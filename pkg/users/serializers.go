package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
} //@name User

type UserSerializer struct {
	C *gin.Context
	UserModel
}

func (s *UserSerializer) Response() UserResponse {

	response := UserResponse{
		ID:       s.ID,
		Username: s.Username,
		Email:    s.Email,
	}

	return response
}

type UsersSerializer struct {
	C     *gin.Context
	Users []UserModel
}

func (s *UsersSerializer) Response() []UserResponse {
	var response []UserResponse
	for _, user := range s.Users {
		serializer := UserSerializer{s.C, user}
		response = append(response, serializer.Response())
	}

	return response
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TokenSerializer struct {
	C     *gin.Context
	Token string
} //@name Token

func (s *TokenSerializer) Response() TokenResponse {
	response := TokenResponse{
		Token: s.Token,
	}

	return response
}
