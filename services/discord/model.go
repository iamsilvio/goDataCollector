package discord

// Message Discord API Message
type Message struct {
	Title string `json:"title"`
	Body  string `json:"description"`
}

// MessageContainer Discord API MessageContainer
type MessageContainer struct {
	Content string  `json:"content"`
	Tts     bool    `json:"tts"`
	Mention bool    `json:"allowed_mentions"`
	Message Message `json:"embed"`
}
