package main

import (
	"github.com/Roh-bot/Golang/messagebrokers"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := messagebrokers.NewRabbitInstance()
	if err != nil {
		return
	}
	messagebrokers.RabbitMq.NewPublisher()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		go messagebrokers.RabbitMq.ReceiveMessages()
	}

	time.Sleep(time.Second * 5)
	messagebrokers.RabbitMq.SendMessages(map[string]any{
		"message": "Hello World",
	})

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
