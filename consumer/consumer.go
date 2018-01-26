package consumer

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func NewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:32768",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func ExampleSub() string {
	pubsub := client.Subscribe("mychannel1")
	defer pubsub.Close()

	// Wait for subscription to be created before publishing message.
	subscr, err := pubsub.ReceiveTimeout(time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(subscr)

	msg, err := pubsub.ReceiveMessage()
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.Channel, msg.Payload)
	return  msg.Payload
}
