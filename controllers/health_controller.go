package controllers

import (
	"net/http"

	"41x3n.yama/emoji-splitter/models"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	message := models.Message{Message: "OK"}
	c.JSON(http.StatusOK, message)
}
