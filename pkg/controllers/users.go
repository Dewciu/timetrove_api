package controllers

import (
	"net/http"

	_ "github.com/dewciu/timetrove_api/docs"

	"github.com/dewciu/timetrove_api/pkg/database/queries"
	"github.com/dewciu/timetrove_api/pkg/serializers"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetAllUsers godoc
// @Summary Ping example
// @Description Retrieves all users from the database
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} serializers.UserResponse "user"
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	users, err := queries.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	serializer := serializers.UsersSerializer{C: c, Users: users}

	c.JSON(http.StatusOK, serializer.Response())
}
