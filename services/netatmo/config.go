package netatmo

// Config Netatwo API configuration
type Config struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}
