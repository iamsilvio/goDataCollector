package discord

type DiscordConfig struct {
	WebHookUrl string `json:"webHookUrl"`
}

type Message struct {
	Title string `json:"title"`
	Body  string `json:"description"`
}

type MessageContainer struct {
	Content string  `json:"content"`
	Tts     bool    `json:"tts"`
	Mention bool    `json:"allowed_mentions"`
	Message Message `json:"embed"`
}
