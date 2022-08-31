package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Multiplication() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printMulTable(i)
	}
	wg.Wait()
}

func printMulTable(num int) {
	for i := 2; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, num, num*i)
		time.Sleep(1 * time.Millisecond)
	}
	wg.Done()
}
