package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()

	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()

	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()

	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()

	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()
	go publish()

	go publish()
	go publish()
	go publish()
	time.Sleep(10)
	go subscribe()
	go subscribe()
	go subscribe()
	select {}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func publish() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//
	for {
		select {
		case <-time.After(time.Second):
			body := "hello"
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			failOnError(err, "Failed to publish a message")

		}
	}
}

func subscribe() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
