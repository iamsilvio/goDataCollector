package wellknownip

import (
	"log"
	"net/http"
)

// GetMyPublicIP returns the string of the Public ip
func GetMyPublicIP() string {

	ip := ""

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

	ip = resp.Header.Get("X-DeleteOnError-IP")

	return ip
}
