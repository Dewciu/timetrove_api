package routes

import (
	_ "github.com/dewciu/timetrove_api/docs"
	"github.com/dewciu/timetrove_api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", controllers.GetAllUsers)
	}
}
