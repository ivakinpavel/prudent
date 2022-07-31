package main

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/malleator/prudent/pkg/app"
	"gitlab.com/malleator/prudent/pkg/config"
)

func main() {
	log.Info("Starting")
	conf, err := config.GetConfig(".env")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Configuration error")
	}
	log.Info("Config loaded successfully")

	app.Start(&conf)
}
