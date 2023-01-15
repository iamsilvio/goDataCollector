package pushover

import (
	"bytes"
	"encoding/json"

	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var config Config

// SetConfig sets the API configuration
func SetConfig(conf Config) {
	config = conf
}

// PushIPChange to pushover api
func PushIPChange(ip string) {

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Failed to Post new Ip: %v\n", r)
		}
	}()

	url := "https://api.pushover.net/1/messages.json"

	data := Message{Token: config.ApplicationKey, User: config.UserKey, Message: "Ip has Changed @Home: " + ip}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", url, reqBodyBytes)
	if err != nil {
		log.WithError(err).Errorf("Could not create POST request\n")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Failed to Post new IP to Pushover\n")
	}
	defer resp.Body.Close()

}

// PushNotification to pushover api
func PushNotification(co2 int) {

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Failed to Post co2 value: %v\n", r)
		}
	}()

	url := "https://api.pushover.net/1/messages.json"

	data := Message{Token: config.ApplicationKey, User: config.UserKey, Message: "Co2 Value: " + strconv.Itoa(co2)}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", url, reqBodyBytes)
	if err != nil {
		log.WithError(err).Errorf("Could not create POST request\n")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("Failed to Post Co2 Warning to Pushover\n")
	}
	defer resp.Body.Close()

}
