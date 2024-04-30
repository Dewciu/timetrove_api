package routes

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func (r routes) addSwaggerRoutes(rg *gin.RouterGroup) {
	swag := rg.Group("/swagger")
	swag.GET("/*any", swagger.WrapHandler(files.Handler))
}
