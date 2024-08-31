package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jasurxaydarov/todo_app_kafka/service/models"
	"github.com/segmentio/kafka-go"
)

var Message []models.Message = []models.Message{}

func NewConsumer(topic string) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"0.0.0.0:9092"},
		GroupID: "github.com/jasurxaydarov/todo_app_kafka",
		Topic:   topic,
	})

	fmt.Println(fmt.Sprintln(topic," consumer is listenning"))
	for {

		var msg models.Message
		Kfmsg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("err on ReadMessage", err)
			continue
		}

		err = json.Unmarshal(Kfmsg.Value, &msg)

		if err != nil {
			log.Println("err on Unmarshal", err)
			continue
		}

		switch topic {
		case "create-message":

			Message=append(Message, msg)

			fmt.Println("new message:",msg)
			fmt.Println("all messages:",Message)


		}
	}

}

func RunConsumers() {

	go NewConsumer("Get-message")

	go NewConsumer("create-message")

}
