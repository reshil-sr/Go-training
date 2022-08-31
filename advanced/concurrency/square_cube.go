package concurrency

import (
	"fmt"
	"time"
)

var intSlice = []int{1, 2, 3, 4}

func SquareCube() {
	fmt.Println("Exercise 11.2 Find Square and Cube from an array using goroutine")
	sqCh := make(chan int, 4)
	cuCh := make(chan int, 4)
	doneCh := make(chan int)

	for _, num := range intSlice {
		go calSquare(num, sqCh)
		go calCube(num, cuCh)
	}

	go prntSq(sqCh, doneCh)
	go prntCb(cuCh, doneCh)

	time.Sleep(1 * time.Millisecond)
}

func calSquare(num int, sq chan int) {
	sq <- num * 2
}

func calCube(num int, cb chan int) {
	cb <- num * 3
}

func prntSq(sq chan int, d chan int) {
	fmt.Println("Square of given numbers are:")

	for i := 0; i <= 4; i++ {
		fmt.Println(<-sq)
	}
	time.Sleep(1 * time.Millisecond)
	d <- 1
}

func prntCb(cb chan int, d <-chan int) {
	fmt.Println("Cube of given numbers are:")
	for i := 0; i <= 4; i++ {
		fmt.Println(<-cb)
	}
}
