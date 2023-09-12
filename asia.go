package main

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"asia",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	
	go func() {
		for d := range msgs {
			log.Println(d.Body)
		}
	}()

	<-forever
}