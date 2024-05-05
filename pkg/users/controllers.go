package users

import (
	"errors"
	"fmt"
	"net/http"

	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/common"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllUsersController godoc
// @Summary Get Users
// @Description Retrieves all users from the database
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse "Returns list of users"
// @Router /users [get]
func GetAllUsersController(c *gin.Context) {
	users, err := GetAllUsersQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
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
// @Success 200 {object} UserResponse "Returns Created User"
// @Router /users [post]
func CreateUserController(c *gin.Context) {
	validator := UserModelValidator{}
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidationError(err))
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
