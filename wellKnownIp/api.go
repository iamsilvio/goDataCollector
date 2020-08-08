package wellKnownIp

import (
	"log"
	"net/http"
)

func GetMyPublicIp() string {
	resp, err := http.Get("http://deleteonerror.com/.well-known/ip")
	if err != nil {
		log.Printf("Failed to get public ip: %v\n", err)
	}

	defer resp.Body.Close()

	ip := resp.Header.Get("X-DeleteOnError-IP")

	return ip
}
