package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/todo_app_kafka/api/api/handlers"
)

func Api(){

	engine:=gin.Default()

	api:=engine.Group("/api")

	api.POST("create-message",handlers.CreateMessage)

	engine.Run()
}