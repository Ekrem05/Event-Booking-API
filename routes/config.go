package routes

import (
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func AddRoutes(server *gin.Engine){
	server.GET("/",index)
	server.POST("events",middlewares.Authenticate,createEvent)
	server.GET("/events/:id",getEvent)
	server.PUT("/events/:id",middlewares.Authenticate,updateEvent)
	server.DELETE("/events/:id",middlewares.Authenticate,deleteEvent)

	server.POST("/events/:id/register",middlewares.Authenticate,register)
	server.DELETE("/events/:id/register",middlewares.Authenticate,remove)

	server.POST("/signup",signUp)
	server.POST("/login",login)



}