package discord

import (
	"bytes"
	"encoding/json"

	"net/http"

	log "github.com/sirupsen/logrus"
)

var config Config

// SetConfig sets the API configuration
func SetConfig(conf Config) {
	config = conf
}

// PushIPChange to discord webhook
func PushIPChange(ip string) {
	log.Trace("Pushing Ip change to Discord\n")

	msg := Message{Title: "Server Ip has Changed", Body: "New Ip: " + ip}

	data := MessageContainer{Content: msg.Title + "\n" + msg.Body, Tts: false, Mention: false, Message: msg}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", config.WebHookURL, reqBodyBytes)
	if err != nil {
		log.WithError(err).Errorf("Could not create POST request\n")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Errorf("Failed to Post new Ip to Discord \n")
	}
	defer resp.Body.Close()

}
