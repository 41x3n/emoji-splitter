package controllers

import (
	"fmt"
	"net/http"

	"41x3n.yama/emoji-splitter/models"
	"41x3n.yama/emoji-splitter/utils"
	"github.com/gin-gonic/gin"
)

var errorLogger = utils.ErrorLogger
var infoLogger = utils.InfoLogger

func SplitEmojiCommon(c *gin.Context, emoji string) (models.EmojiResponse, error) {
	if emoji == "" {
		return models.EmojiResponse{}, fmt.Errorf("invalid emoji received")
	}

	characters := utils.ExtractEmojiCharacters(emoji)
	if len(characters) == 0 {
		return models.EmojiResponse{}, fmt.Errorf("no emoji found in payload: %s", emoji)
	}

	response := models.EmojiResponse{
		OriginalMessage: emoji,
		SplitEmojiList:  characters,
		Message:         "We have trimmed the payload to return only emoji characters",
	}

	infoLogger.Printf("Emoji Response: %+v", response)
	return response, nil
}

func SplitEmojiPost(c *gin.Context) {
	var payload models.EmojiPayload
	if err := c.BindJSON(&payload); err != nil {
		errorLogger.Printf("Error binding JSON: %s", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	response, err := SplitEmojiCommon(c, payload.Emoji)
	if err != nil {
		errorLogger.Printf("%s", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func SplitEmojiGet(c *gin.Context) {
	emoji := c.Query("emoji")

	response, err := SplitEmojiCommon(c, emoji)
	if err != nil {
		errorLogger.Printf("%s", err.Error())
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
