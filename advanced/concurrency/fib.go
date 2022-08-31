package concurrency

import (
	"fmt"
	"time"
)

func Fib() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
		quit <- false
	}(n)

	go fibonacci(ch, quit)
	time.Sleep(1 * time.Millisecond)
}

func fibonacci(ch chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
