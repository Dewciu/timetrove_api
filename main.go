package main

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/config"
	"github.com/dewciu/timetrove_api/pkg/database"
	"github.com/dewciu/timetrove_api/pkg/routes"
	"github.com/sirupsen/logrus"

	_ "github.com/dewciu/timetrove_api/pkg/docs"
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

	router := routes.SetupRouter()

	if err = database.Connect(conf); err != nil {
		msg := fmt.Sprintf("Failed to connect to database: %v", err)
		panic(msg)
	}

	defer database.Disconnect()

	if err = database.Migrate(); err != nil {
		msg := fmt.Sprintf("Failed to migrate database: %v", err)
		panic(msg)
	}

	hostname := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	router.Run(hostname)
}
