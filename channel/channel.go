package channel

// 消费者
func Consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv: ", x)
	}

	done <- true
}

// 生产者
func Producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i // 发送数据
	}

	close(data)
}
