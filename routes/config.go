package routes

import "github.com/gin-gonic/gin"

func AddRoutes(server *gin.Engine){
	server.GET("/",index)
	server.POST("events",createEvent)
	server.GET("events/:id",getEvent)
	server.PUT("events/:id",updateEvent)
	server.DELETE("events/:id",deleteEvent)

}