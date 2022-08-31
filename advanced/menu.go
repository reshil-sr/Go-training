package advanced

import (
	"fmt"
	"gotraining/advanced/concurrency"
	"gotraining/advanced/polymorphism"
)

func Menu() {
	var c float32

	fmt.Println("Choose Advanced Exercises: ")
	fmt.Println("Enter 10.1 for Exercise 10 Polymorphism: ")
	fmt.Println("Enter 11.1 for Exercise 11.1 Find Odd and Even numbers from an array using goroutine: ")
	fmt.Println("Enter 11.2 for Exercise 11.2 Find Square and Cube of given numbers using goroutine: ")
	fmt.Println("Enter 11.3 for Exercise 11.3 Read & Write Fibonacci using goroutine: ")
	fmt.Println("Enter 11.5 for Exercise 11.5 Multiplication table using concurrency: ")
	fmt.Println("Enter 11.6 for Exercise 11.6 Bug Fix: ")

	fmt.Scanf("%f", &c)

	switch c {
	case 10.1:
		// Exercise 10 Polymorphism
		polymorphism.ReadTime()
	case 11.1:
		// Exercise 11.1 Polymorphism
		// concurrency.OddEven()
		concurrency.EvenOdd()
	case 11.2:
		// Exercise 11.2 Square & Cube
		concurrency.SquareCube()
	case 11.3:
		// Exercise 11.3 Fibonacci
		concurrency.Fib()
	case 11.5:
		// Exercise 11.5 Multiplication
		concurrency.Multiplication()
	case 11.6:
		// Exercise 11.6 BugFix
		concurrency.BugFix()
	default:
		fmt.Println("Invalid Choice")
	}
}
