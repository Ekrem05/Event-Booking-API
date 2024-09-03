package middlewares

import (
	"api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token:=context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	userId,err:=utils.Verify(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	context.Set("userId",userId)
	context.Next()
}