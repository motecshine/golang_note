package concurrency

import (
	"fmt"
	"time"
)

func Task(id int) {
	for i := 1; i < 5; i++ {
		fmt.Printf("%d：%d\n", id, i)
		time.Sleep(time.Second)
	}
}
