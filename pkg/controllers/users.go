package controllers

import (
	"net/http"

	_ "github.com/dewciu/timetrove_api/pkg/docs"

	"github.com/dewciu/timetrove_api/pkg/database"
	"github.com/dewciu/timetrove_api/pkg/database/models"
	"github.com/gin-gonic/gin"
)

// @Summary Ping example
// @Description This is a ping example
// @scheme http
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
