package users

import (
	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func AddUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", middleware.AuthMiddleware(), GetAllUsersController)
		users.POST("/", middleware.AuthMiddleware(), CreateUserController)
		users.GET("/:id", middleware.AuthMiddleware(), GetUserByIDController)
		users.DELETE("/:id", middleware.AuthMiddleware(), DeleteUserByIDController)
		users.PUT("/:id", middleware.AuthMiddleware(), UpdateUserController)
	}
}

func AddAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", LoginController)
	}
}
