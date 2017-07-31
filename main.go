package main

import (
	"fmt"

	"./channel"
)

func main() {
	// 同步阻塞并发
	done := make(chan bool)         //接受消费结束的信号
	data := make(chan int)          //数据管道
	go channel.Consumer(data, done) // 启动消费者
	go channel.Producer(data)       // 启动生产者

	<-done // 阻塞， 知道消费者发回结束信号
	fmt.Print("hh")
}
