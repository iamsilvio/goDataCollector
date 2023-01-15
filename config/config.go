package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func ReadConfig(path string) DataCollectorConfig {

	data, err := ioutil.ReadFile(path)
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
