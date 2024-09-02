package routes

import (
	"api/models"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context){

	var user models.User

	err:=context.ShouldBindJSON(&user)

	if err!=nil{
		context.JSON(400,gin.H{"error":"Invalid body"})
		return;
	}

	err=user.Save()

	if err!=nil{
		context.JSON(500,gin.H{"error":"Could not save the user"})
		return;
	}

	context.JSON(200,nil)
	
	
}