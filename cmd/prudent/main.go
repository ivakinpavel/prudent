package main

import (
	"github.com/ivakinpavel/prudent/pkg/app"
	"github.com/ivakinpavel/prudent/pkg/config"
	log "github.com/sirupsen/logrus"
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
