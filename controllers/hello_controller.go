package controllers

import (
	"net/http"

	"41x3n.yama/emoji-splitter/models"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	message := models.Message{Message: "Hello, World!"}
	c.JSON(http.StatusOK, message)
}
