package netAtmo

import (
	"context"
	"encoding/json"
	"log"

	"golang.org/x/oauth2"
)

var token *oauth2.Token
var apiURL string
var oauthConfig *oauth2.Config

var ctx = context.Background()

func SetConfig(conf ApiConfig) {

	apiConfig := conf
	baseURI := "https://api.netatmo.com/"

	apiURL = baseURI + "api/"
	netatmoEndpoint := oauth2.Endpoint{
		AuthURL:  baseURI + "oauth2/authorize",
		TokenURL: baseURI + "oauth2/token",
	}

	oauthConfig = &oauth2.Config{
		ClientID:     apiConfig.ClientID,
		ClientSecret: apiConfig.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint:     netatmoEndpoint,
	}

	var err error

	token, err = oauthConfig.PasswordCredentialsToken(ctx, apiConfig.UserName, apiConfig.Password)
	if err != nil {
		log.Printf("Failed Pwd cred token: %v\n", err)

	}
}

func GetStationsData() (Dashboard, error) {

	var result Dashboard

	defer func() {
		if r := recover(); r != nil {
			result = Dashboard{}
		}
	}()

	client := oauthConfig.Client(ctx, token)
	var dev ApiResponse

	resp, err := client.Get(apiURL + "getstationsdata")
	if err != nil {
		log.Printf("Failed to get station data: %v\n", err)
		return dev.Body.Devices[0].Dashboard, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&dev)
	if err != nil {
		log.Printf("Failed to parse Netatmo Response: %v\n", err)
	}

	result = dev.Body.Devices[0].Dashboard

	return result, nil
}
