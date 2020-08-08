package pushOver

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var config PushOverConfig

func SetConfig(conf PushOverConfig) {
	config = conf
}

func PushIpChange(ip string) {

	url := "https://api.pushover.net/1/messages.json"

	data := Message{Token: config.ApplicationKey, User: config.UserKey, Message: "Ip has Changed @Home: " + ip}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	req, err := http.NewRequest("POST", url, reqBodyBytes)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to Post new Ip to Discord: %v\n", err)
	}
	defer resp.Body.Close()

	//	fmt.Println("response Status:", resp.Status)
	//	fmt.Println("response Headers:", resp.Header)
	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println("response Body:", string(body))

}
