package main

import "github.com/jasurxaydarov/todo_app_kafka/service/kafka"


func main(){

	kafka.RunConsumers()

	select{} 
}