package controller

import "github.com/gin-gonic/gin"

type helloWorld struct {
}

func (helloWorld) Welcome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "Hello World",
	})
}

// HelloWorld : instance to helloworld struct
var HelloWorld helloWorld
