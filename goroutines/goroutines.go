package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Loop(Channel chan int) {
	var m *sync.Mutex
	m = new(sync.Mutex)
	m.Lock()
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		fmt.Printf("%d\n", i)
	}
	m.Unlock()
	Channel <- 0 // 数据流入
}
