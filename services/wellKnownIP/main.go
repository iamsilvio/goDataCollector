package wellKnownIP

import (
	"encoding/json"
	"os"

	"net/http"

	log "github.com/sirupsen/logrus"
)

var lastIP = ""
var dataPath = ""

func SetConfig(path string) {
	dataPath = path
}

// GetMyPublicIP returns the string of the Public ip
func GetMyPublicIP() (ip string, hasChanged bool) {
	log.Trace("GetMyPublicIP\n")

	if len(lastIP) == 0 {
		LoadIP()
	}

	defer func() {
		if r := recover(); r != nil {
			ip = ""
		}
	}()

	resp, err := http.Get("http://49.12.2.55/.well-known/ip")
	if err != nil {
		log.Printf("Failed to get public ip: %v\n", err)
	}

	defer resp.Body.Close()

	tmpIp := resp.Header.Get("X-DeleteOnError-IP")

	if len(lastIP) == 0 {
		log.Infof("Public IP was empty and has to be saved %s\n", tmpIp)
		hasChanged = false
		ip = tmpIp
		lastIP = tmpIp
		SaveIP()
		return
	}

	if tmpIp != lastIP {
		log.Infof("Public IP has changed %s\n", tmpIp)
		hasChanged = true
		ip = tmpIp
		lastIP = tmpIp
		SaveIP()

	}
	return
}

func SaveIP() {
	log.Tracef("Saving %s\n", lastIP)
	file, err := json.MarshalIndent(lastIP, "", " ")
	if err != nil {
		log.WithError(err).Error("Failed to marshal ip\n")
		return
	}

	err = os.WriteFile(dataPath+"ip.json", file, 0644)
	if err != nil {
		log.WithError(err).Error("Failed to write ip file\n")
	}
}

func LoadIP() {
	log.Trace("Load IP \n")

	file, err := os.ReadFile(dataPath + "ip.json")
	if err != nil {
		log.WithError(err).Error("Failed to read ip file\n")
		lastIP = ""
		return
	}

	err = json.Unmarshal([]byte(file), &lastIP)
	if err != nil {
		log.WithError(err).Error("Failed to parse ip file\n")
		lastIP = ""
	}

}
