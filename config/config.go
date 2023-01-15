package config

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

func ReadConfig(path string) DataCollectorConfig {

	data, err := os.ReadFile(path)
	if err != nil {
		log.WithError(err).Error("Failed to read config file\n")
	}

	var conf DataCollectorConfig

	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.WithError(err).Error("Failed to parse config file\n")
	}
	return conf
}
