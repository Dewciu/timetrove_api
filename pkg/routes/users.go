package routes

import (
	"github.com/dewciu/timetrove_api/pkg/controllers"
	_ "github.com/dewciu/timetrove_api/pkg/docs"
	"github.com/gin-gonic/gin"
)

func (r routes) addUsersRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", controllers.GetAllUsers)
	}
}
