package main

import (
	"fmt"
	"time"
)

var msgchan = make(chan string)

func worker() {
}

func main() {
	go func() {

		for i := 0; i < 10; i++ {
			go func(workid int) {
				for {
					select {
					case msg := <-msgchan:
						fmt.Printf("msg:%s \n", msg)
					}
				}
			}(i)
		}
	}()

	for i := 0; i < 1000000; i++ {
		msgchan <- fmt.Sprintf("msg in %d", i)
	}
	time.Sleep(time.Hour)
}
