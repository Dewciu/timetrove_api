package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func SetupRouter() *gin.Engine {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/api/v1")
	r.addSwaggerRoutes(v1)
	r.addUsersRoutes(v1)
	return r.router
}
