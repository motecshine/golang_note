package main

import (
    "flag"
	"golang_note/websocket"
	"log"
	"net/http"
	"golang_note/consumer"
)

func main() {
	// // 同步阻塞并发
	// done := make(chan bool)         //接受消费结束的信号
	// data := make(chan int)          //数据管道
	// go channel.Consumer(data, done) // 启动消费者
	// go channel.Producer(data)       // 启动生产者

	// <-done // 阻塞， 知道消费者发回结束信号

	// goroutine
	// var Channel chan int = make(chan int)

	// go goroutines.Loop(Channel)
	// <-Channel

	// go goroutines.Loop(Channel)

	// <-Channel

	consumer.NewClient()

	consumer.ExampleSub()
	var addr = flag.String("addr", "localhost:8080", "http service address")
	http.HandleFunc("/echo", websocket.Server)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
