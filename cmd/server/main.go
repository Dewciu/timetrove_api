package main

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/addresses"
	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/dewciu/timetrove_api/pkg/config"
	"github.com/dewciu/timetrove_api/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	_ "github.com/dewciu/timetrove_api/docs"
)

// @title TimeTrove API
// @version 1.0
// @description This is an API for TimeTrove application
// @termsOfService http://swagger.io/terms/

// @contact.name Kacper Król
// @contact.email kacperkrol99@icloud.com

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @Schemes http https
func main() {
	conf, err := config.GetConfig()

	if err != nil {
		logrus.Panicf("Failed to get configuration: %v", err)
	}

	router := SetupRouter()

	if err = common.Connect(conf); err != nil {
		msg := fmt.Sprintf("Failed to connect to common: %v", err)
		panic(msg)
	}

	defer common.Disconnect()

	if err = Migrate(); err != nil {
		msg := fmt.Sprintf("Failed to migrate common: %v", err)
		panic(msg)
	}

	hostname := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	router.Run(hostname)
}

func Migrate() error {
	if err := common.DB.AutoMigrate(
		&users.UserModel{},
		&addresses.AddressModel{},
	); err != nil {
		return err
	}

	return nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	addSwaggerRoutes(v1)
	users.AddUsersRoutes(v1)
	return r
}

func addSwaggerRoutes(rg *gin.RouterGroup) {
	swag := rg.Group("/swagger")
	swag.GET("/*any", swagger.WrapHandler(files.Handler))
}