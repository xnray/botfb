package model

type SendMessage struct {
	MessagingType string      `json:"messaging_type"`
	Recipient     Recipient   `json:"recipient"`
	Message       MessageData `json:"message"`
	Tag           string      `json:"tag,omitempty"`
}

type Recipient struct {
	ID int64 `json:"id,string"`
}

type MessageData struct {
	Text string `json:"text,omitempty"`
}
