package concurrency

import (
	"fmt"
	"time"
)

//first approach works but required to add timeout to print all values since main go routine exits before other go routines complete execution
func OddEven() {
	fmt.Println("Exercise 11.1 Find Odd and Even numbers from an array using goroutine")
	intSlice := []int{1, 2, 3, 4, 33, 46, 54, 55, 77, 66, 99}

	oddChn := make(chan int)
	evenChn := make(chan int)

	go prntEven(evenChn)
	go prntOdd(oddChn)

	for _, num := range intSlice {
		if num%2 == 0 {
			evenChn <- num
		} else {
			oddChn <- num
		}
	}
	time.Sleep(1 * time.Millisecond)
}

func prntEven(c <-chan int) {
	for num := range c {
		fmt.Println("EVEN :", num)
	}
}

func prntOdd(c <-chan int) {
	for num := range c {
		fmt.Println("ODD :", num)
	}
}

//Second Approach this seems a better solution
func EvenOdd() {
	fmt.Println("Exercise 11.1 Find Odd and Even numbers from an array using goroutine")
	intSlice := []int{1, 2, 3, 4, 33, 46, 54, 55, 77, 66, 99}
	numCh := make(chan int)

	go chkEvenOdd(numCh)

	for _, num := range intSlice {
		numCh <- num
	}
}

func chkEvenOdd(c <-chan int) {
	for num := range c {
		if num%2 == 0 {
			fmt.Println("EVEN :", num)
		} else {
			fmt.Println("ODD :", num)
		}
	}
}
