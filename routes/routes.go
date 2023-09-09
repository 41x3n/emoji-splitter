package routes

import (
	"41x3n.yama/emoji-splitter/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/hello", controllers.Hello)

	router.GET("/health", controllers.Health)

	router.POST("/split-emoji", controllers.SplitEmoji)
}
