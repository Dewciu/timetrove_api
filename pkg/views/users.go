package views

import (
	"net/http"

	"github.com/dewciu/timetrove_api/pkg/database"
	"github.com/dewciu/timetrove_api/pkg/database/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
