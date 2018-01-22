package websocket

import (
	"log"
	"net/http"

	"golang_note/consumer"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(*http.Request) bool {
		return true
	},
} // use default options

func Server(w http.ResponseWriter, r *http.Request) {
	consumer.NewClient()

	sub := consumer.Client.Subscribe("mychannel1")
	defer sub.Close()

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		err = c.WriteMessage(1, []byte(msg.Payload))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
