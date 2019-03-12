package router

import (
	"helloWorldGoTest/app/controller"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
)

// Initialize : register all routes and run server
func Initialize() {
	router := gin.New()
	routGroup := router.Group("/")
	registerPublicRoutes(routGroup)

	graceful.Run("0.0.0.0:8000", 120*time.Second, router)
}

func registerPublicRoutes(router *gin.RouterGroup) {
	router.GET("/helloWorld", controller.HelloWorld.Welcome)
}
