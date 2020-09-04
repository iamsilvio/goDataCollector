package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var config Config

// SetConfig sets the API configuration
func SetConfig(conf Config) {
	config = conf
}

// PushIPChange to discord webhook
func PushIPChange(ip string) {

	msg := Message{Title: "7Days Server Ip has Changed", Body: "New Ip: " + ip}

	data := MessageContainer{Content: msg.Title + "\n" + msg.Body, Tts: false, Mention: false, Message: msg}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", config.WebHookURL, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to Post new Ip to Discord: %v", err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

}
