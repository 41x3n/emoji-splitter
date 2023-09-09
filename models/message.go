package models

type Message struct {
	Message string `json:"message"`
}

type EmojiResponse struct {
	OriginalMessage string   `json:"original_message"`
	SplitEmojiList  []string `json:"split_emoji_list"`
	Message         string   `json:"message"`
}
