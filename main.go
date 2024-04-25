package main

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/config"
	"github.com/dewciu/timetrove_api/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		logrus.Panicf("Failed to get configuration: %v", err)
	}

	router := gin.Default()

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
