package pushover

// Config API access config
type Config struct {
	UserKey        string `json:"userKey"`
	ApplicationKey string `json:"appKey"`
}

// Message object of Pushover API
type Message struct {
	Token   string `json:"token"` // your application's API token
	User    string `json:"user"`  // the user/group key
	Message string `json:"message"`
	// Optional parameters
	Device   string `json:"device"`    // your user's device name to send the message (multiple devices may be separated by a comma)
	Title    string `json:"title"`     // title, otherwise your app's name is used
	URL      string `json:"url"`       // - a supplementary URL to show with your message
	URLTitle string `json:"url_title"` // - a title for your supplementary URL
	Priority string `json:"priority"`  // - send as -2 to generate no notification/alert, -1 to always send as a quiet notification, 1 to display as high-priority and bypass the user's quiet hours, or 2 to also require confirmation from the user
	Sound    string `json:"sound"`     // - the name of one of the sounds supported by device clients to override the user's default sound choice

}
