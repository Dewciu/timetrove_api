package main

import (
	"github.com/dewciu/timetrove_api/pkg/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.GetConfiguration()

	if err != nil {
		log.Errorf("Failed to get configuration: %v", err)
		panic("Failed to get configuration!")
	}

	log.Printf("Configuration: %v", conf)
}
