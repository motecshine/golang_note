package main

import (
	"fmt"
	"golang_note/hashtable"
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

	// consumer.NewClient()

	// consumer.ExampleSub()
	// var addr = flag.String("addr", "localhost:8080", "http service address")
	// http.HandleFunc("/echo", websocket.Server)
	// log.Fatal(http.ListenAndServe(*addr, nil))
	// Link := link.InitStack()
	// Link = link.Push(4, Link)
	// Link = link.Push(5, Link)
	// link.PrintLink(Link)
	// Link = link.Pop(Link)
	// link.PrintLink(Link)
	// Link = link.Push(7, Link)
	// Link = link.Push(8, Link)
	// link.PrintLink(Link)*/
	h := new(hashtable.Hashtable)
	h.Add("test", "value")
	h.Add("test2", "value1")
	h.Update("test2", "value2")
	h.Add("hhh", "哈哈哈啊")
	h.Add("hhh", "哈哈s啊")
	fmt.Println(h.Get("test2"), h.Get("test"), h.Get("hhh"))
	h.Delete("hhh")
	fmt.Println(h.Get("test2"), h.Get("test"), h.Get("hhh"))
}
