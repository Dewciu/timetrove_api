package users

import (
	"github.com/dewciu/timetrove_api/pkg/permissions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
} //@name UserResponse

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

type UserWithPermissionsResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Permissions []permissions.PermissionModel
} //@name UserWithPermissionsResponse

type PermissionsForUserSerializer struct {
	C *gin.Context
	UserModel
	Permissions []permissions.PermissionModel
}

func (s *PermissionsForUserSerializer) Response() []permissions.PermissionModel {
	return s.Permissions
}

type TokenResponse struct {
	Token string `json:"token"`
} //@name TokenResponse

type TokenSerializer struct {
	C     *gin.Context
	Token string
}

func (s *TokenSerializer) Response() TokenResponse {
	response := TokenResponse{
		Token: s.Token,
	}

	return response
}
