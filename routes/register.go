package routes

import (
	"api/models"
	events "api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func register(context *gin.Context){
	id,err:=strconv.ParseInt(context.Param("id"),10,64);
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}

	_, err = events.GetById(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "This event does not exist"})
		return
	}
//check if already registered

	registration:=models.Registration{
		User_id:context.GetInt64("userId") ,
		Event_id: id,
	}
	
	err=registration.New();
}

func remove(context *gin.Context){
	id,err:=strconv.ParseInt(context.Param("id"),10,64);
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}

	_, err = events.GetById(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "This event does not exist"})
		return
	}
//check if already registered

	registration:=models.Registration{
		User_id:context.GetInt64("userId") ,
		Event_id: id,
	}
	
	err=registration.Remove();
}