package pushover

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
			log.Printf("Failed to Post new Ip: %v\n", r)
		}
	}()

	url := "https://api.pushover.net/1/messages.json"

	data := Message{Token: config.ApplicationKey, User: config.UserKey, Message: "Ip has Changed @Home: " + ip}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", url, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to Post new Ip to Pushover: %v\n", err)
	}
	defer resp.Body.Close()

	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println("response Body:", string(body))

}

// PushNotification to pushover api
func PushNotification(co2 int) {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Failed to Post co2 value: %v\n", r)
		}
	}()

	url := "https://api.pushover.net/1/messages.json"

	data := Message{Token: config.ApplicationKey, User: config.UserKey, Message: "Co2 Value: " + strconv.Itoa(co2)}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", url, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to Post Co2 Warning to Pushover: %v\n", err)
	}
	defer resp.Body.Close()

}
