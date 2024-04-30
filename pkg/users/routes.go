package users

import (
	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", GetAllUsersController)
		users.POST("/", CreateUserController)
	}
}
