package consumer

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func NewClient() {
<<<<<<< HEAD
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:32768",
=======
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
>>>>>>> 39ae4714135b3fc844f7b9eeafd9828ed1c8ce50
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}

<<<<<<< HEAD
func ExampleSub() string {
	pubsub := client.Subscribe("mychannel1")
=======
func ExampleSub() {
	pubsub := Client.Subscribe("mychannel1")
>>>>>>> 39ae4714135b3fc844f7b9eeafd9828ed1c8ce50
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
