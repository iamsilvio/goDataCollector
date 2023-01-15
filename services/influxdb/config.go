package influxdb

// Config of the influxDb connection
type Config struct {
	ServerURL    string `json:"serverurl"`
	Bucket       string `json:"bucket"`
	Token        string `json:"token"`
	Organisation string `json:"organisation"`
}
