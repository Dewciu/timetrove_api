package users

import (
	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.POST("/login", LoginController)
		users.GET("/", middleware.AuthMiddleware(), GetAllUsersController)
		users.POST("/", middleware.AuthMiddleware(), CreateUserController)
		users.GET("/:id", middleware.AuthMiddleware(), GetUserByIDController)
	}
}
