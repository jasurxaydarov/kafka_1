package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/todo_app_kafka/api/kafka"
	"github.com/jasurxaydarov/todo_app_kafka/api/models"
)


func CreateMessage(ctx *gin.Context){

	var req models.Message

	err:=ctx.BindJSON(&req)

	if err!=nil{
		log.Fatalln("err on BindJSON",err)
	}

	err=kafka.CreateMessageProducer(req)

	if err!=nil{

		log.Fatalln("err on CreateMessageProducer", err)

	}

	ctx.JSON(201,"created")
}