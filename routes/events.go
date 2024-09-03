package routes

import (
	events "api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func index(context *gin.Context) {
	events, err := events.GetAll()
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to fetch events"})
		return

	}

	context.JSON(200, events)
}
func createEvent(context *gin.Context) {
	
	var event *events.Event
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event.UserId = context.GetInt64("userId");
	event.DateTime = time.Now()
	err=event.Save()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, event)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}
	event, err := events.GetById(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to get an event by the specified id"})
		return

	}
	context.JSON(200, event)
}

func updateEvent(context *gin.Context){
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}

	event, err := events.GetById(id)

	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to get an event by the specified id"})
		return
	}

	userId:=context.GetInt64("userId");

	if userId!=event.UserId{
		context.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var updatedEvent events.Event 
	err=context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updatedEvent.Id = id;
	err=events.UpdateEvent(&updatedEvent)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200,nil)
}

func deleteEvent(context *gin.Context){
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid parameter"})
		return
	}

	event, err := events.GetById(id)
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to get an event by the specified id"})
		return
	}

	userId:=context.GetInt64("userId");
	if userId!=event.UserId{
		context.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err = events.DeleteEvent(id);
	if err != nil {
		context.JSON(500, gin.H{"error": "Could not delete the event"})
		return

	}
	context.JSON(200, nil)
}