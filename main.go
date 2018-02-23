package main

import (
	"context"
	"fmt"
	c "golang_note/context"
	"log"
	"net/http"
	"time"
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
	// h := new(hashtable.Hashtable)
	// h.Add("test", "value")
	// h.Add("test2", "value1")
	// h.Update("test2", "value2")
	// h.Add("hhh", "哈哈哈啊")
	// h.Add("hhh", "哈哈s啊")
	// fmt.Println(h.Get("test2"), h.Get("test"), h.Get("hhh"))
	// h.Delete("hhh")
	// fmt.Println(h.Get("test2"), h.Get("test"), h.Get("hhh"))
	// newTree := new(tree.TreeController)
	// newTree.Init(4)
	// newTree.Insert(50)
	// newTree.Insert(5)
	// newTree.Insert(7)
	// newTree.Insert(53)
	// newTree.Insert(52)
	// newTree.Insert(56)
	// newTree.Insert(3)

	// 启动一个http 协程
	go http.ListenAndServe(":8989", nil)
	// 获取当前的ctx
	ctx, cancel := context.WithCancel(context.Background())

	// 模拟新协程 新协程将在三秒后退出
	go func() {
		// 睡眠三秒
		time.Sleep(3 * time.Second)
		fmt.Println("阻塞结束")
		// 取消ctx
		cancel()
	}()

	log.Println(c.A(ctx))
	select {}
}
