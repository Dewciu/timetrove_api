package main

import (
	"fmt"
	"net/url"

	"github.com/dewciu/timetrove_api/pkg/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf, err := config.GetConfig()

	if err != nil {
		log.Panicf("Failed to get configuration: %v", err)
	}

	router := gin.Default()

	dsn := url.URL{
		User:     url.UserPassword(conf.Database.User, conf.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", conf.Database.Host, conf.Database.Port),
		Path:     conf.Database.Name,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn.String()}), &gorm.Config{})

	defer func() {
		dbInstance, _ := db.DB()
		if err := dbInstance.Close(); err != nil {
			log.Errorf("Failed to close database connection: %v", err)
		}
	}()

	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	// Configure database models and migrations here
	router.Run(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port))
}
