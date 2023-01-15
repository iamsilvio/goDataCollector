package netatmo

import (
	"context"
	"encoding/json"

	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var apiURL string
var client *http.Client

// SetConfig sets the API configuration
func SetConfig(conf Config) {

	baseURI := "https://api.netatmo.com/"

	apiURL = baseURI + "api/"

	oAuthConfig := &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Scopes:       []string{"read_station"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  baseURI + "oauth2/authorize",
			TokenURL: baseURI + "oauth2/token",
		},
	}

	token := &oauth2.Token{
		AccessToken:  conf.AccessToken,
		RefreshToken: conf.RefreshToken,
	}

	client = oAuthConfig.Client(context.Background(), token)
}

func GetStationsData() (Dashboard, error) {

	var result Dashboard

	defer func() {
		if r := recover(); r != nil {
			result = Dashboard{}
		}
	}()

	var dev APIResponse

	resp, err := client.Get(apiURL + "getstationsdata")
	if err != nil {
		log.WithError(err).Error("Failed to get station data\n")

		return dev.Body.Devices[0].Dashboard, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&dev)
	if err != nil {
		log.WithError(err).Error("Failed to parse Netatmo Response\n")
		return result, err
	}

	result = dev.Body.Devices[0].Dashboard

	log.Debugf("%v\n", dev)

	return result, nil

}

func GetToken() (accesToken string, refreshToken string) {

	token, err := client.Transport.(*oauth2.Transport).Source.Token()
	if err != nil {
		log.WithError(err).Error("Error getting token")
		return
	}
	log.Infof("Token will expire at %s", token.Expiry.Format("01-02-2006 15:04:05"))

	accesToken = token.AccessToken
	refreshToken = token.RefreshToken

	return
}
