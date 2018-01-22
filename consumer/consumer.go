package consumer

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func NewClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}

func ExampleSub() {
	pubsub := Client.Subscribe("mychannel1")
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
}
