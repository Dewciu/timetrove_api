package users

import (
	"errors"
	"fmt"
	"net/http"

	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/common"

	"github.com/gin-gonic/gin"
)

//TODO: Improve error handling -> more descriptive error messages

// @BasePath /api/v1

// LoginController godoc
// @Summary Retrieve JWT API token
// @Description Retrieve JWT API token, when given valid username and password
// @Tags user
// @Accept json
// @Produce json
// @Param Credentials body LoginValidator true "Login Credentials"
// @Success 200 {object} TokenResponse "Returns JWT token"
// @Router /users/login [post]
func LoginController(c *gin.Context) {
	var validator LoginValidator

	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("login", err))
		return
	}

	u := UserModel{Username: validator.Username, Password: validator.Password}

	token, err := u.LoginCheck()

	serializer := TokenSerializer{C: c, Token: token}

	if err != nil {
		c.JSON(http.StatusUnauthorized, common.NewError("login", errors.New("invalid credentials")))
		return
	}

	c.JSON(http.StatusOK, serializer.Response())
}

// @BasePath /api/v1

// GetAllUsersController godoc
// @Summary Get Users
// @Description Retrieves all users from the database
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param email query string false "User's E-mail"
// @Param username query string false "User's username"
// @Param id query string false "User's ID"
// @Success 200 {array} UserResponse "Returns list of users"
// @Router /users [get]
func GetAllUsersController(c *gin.Context) {
	var users []UserModel
	var err error

	if len(c.Request.URL.Query()) == 0 {
		fmt.Println(c.Request.URL.Query())
		users, err = GetAllUsersQuery()
	} else {
		users, err = GetUsersByFilterQuery(c)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}

	serializer := UsersSerializer{C: c, Users: users}

	c.JSON(http.StatusOK, serializer.Response())
}

// @BasePath /api/v1

// CreateUserController godoc
// @Summary Create User
// @Description Creates single user in database
// @Tags user
// @Accept json
// @Produce json
// @Param User body UserModelValidator true "User Object"
// @Security ApiKeyAuth
// @Success 200 {object} UserResponse "Returns Created User"
// @Router /users [post]
func CreateUserController(c *gin.Context) {
	validator := UserModelValidator{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("user", err))
		return
	}

	err := CreateUserQuery(validator.userModel)

	if err != nil {
		fmt.Println(err)
		var er *common.AlreadyExistsError
		if errors.As(err, &er) {
			err := err.(*common.AlreadyExistsError)
			c.JSON(http.StatusConflict, common.NewError("user", err))
			return
		}
		c.JSON(http.StatusInternalServerError, common.NewError("database", err))
		return
	}

	serializer := UserSerializer{C: c, UserModel: validator.userModel}

	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
