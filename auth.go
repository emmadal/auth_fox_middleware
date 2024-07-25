package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var context *gin.Context

func Authenticate() {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	//  validate the token and get user data from it
	userID, err := VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// Attach the userID to the request
	context.Set("userID", userID)

	context.Next()

}
