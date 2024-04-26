package routes

import (
	"github.com/dewciu/timetrove_api/pkg/views"
	"github.com/gin-gonic/gin"
)

func (r routes) addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		// users.POST("/", r.create)
		users.GET("/", views.GetAllUsers)
		// users.GET("/:id", r.retrieve)
		// users.PUT("/:id", r.update)
		// users.DELETE("/:id", r.delete)
	}
}
