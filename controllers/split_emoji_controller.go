package controllers

import (
	"net/http"

	"41x3n.yama/emoji-splitter/models"
	"41x3n.yama/emoji-splitter/utils"
	"github.com/gin-gonic/gin"
)

var errorLogger = utils.ErrorLogger
var infoLogger = utils.InfoLogger

func SplitEmoji(c *gin.Context) {
	var payload models.EmojiPayload

	if err := c.BindJSON(&payload); err != nil {
		errorLogger.Printf("Error binding JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	emoji := payload.Emoji
	infoLogger.Printf("Emoji Payload: %s", emoji)

	if emoji == "" || len(emoji) == 0 {
		errorMessage := "Invalid emoji received"
		errorLogger.Printf("%s", errorMessage)

		c.JSON(http.StatusBadRequest, models.Error{Error: errorMessage})
		return
	}

	var characters []string
	for _, r := range emoji {
		infoLogger.Printf("%c -> %U\n", r, r)
		if utils.IsEmoji(r) {
			characters = append(characters, string(r))
		}
	}

	if len(characters) == 0 {
		errorMessage := "No emoji found in payload"
		errorLogger.Printf("%s: %s", errorMessage, emoji)

		c.JSON(http.StatusBadRequest, models.Error{Error: errorMessage})
		return
	}

	EmojiResponse := models.EmojiResponse{
		OriginalMessage: emoji,
		SplitEmojiList:  characters,
		Message:         "We have trimmed the payload to return only emoji characters",
	}

	infoLogger.Printf("Emoji Response: %+v", EmojiResponse)

	c.JSON(http.StatusOK, EmojiResponse)
}
