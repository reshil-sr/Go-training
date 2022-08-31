package main

import (
	"fmt"
	"gotraining/advanced"
	"gotraining/basic"
	"gotraining/medium"
)

func main() {
	menu()
}

func exBasic() {

	var c int

	fmt.Println("Choose Exercise: ")
	fmt.Println("Enter 1 for Exercise 1.1 Hello world: ")
	fmt.Println("Enter 2 for Exercise 3.1 : Modulo operation: ")
	fmt.Println("Enter 3 for Exercise 3.2 : Modulo operation: ")
	fmt.Println("Enter 4 for EXERCISE 4.1: Multidimensional Slice: ")
	fmt.Println("Enter 5 for EXERCISE 4.2: Slice of int: ")
	fmt.Println("Enter 6 for EXERCISE 4.3:Delete Slice of int: ")
	fmt.Println("Enter 7 for EXERCISE 4.4 Dynamic slice & convert slice to Map: ")
	fmt.Println("Enter 8 for EXERCISE 4.4 Create a menu with add,delete option for slice & map: ")
	fmt.Println("Enter 9 for EXERCISE 4.5 String Map: ")
	fmt.Println("Enter 10 for EXERCISE 4.6 Add to String Map: ")
	fmt.Println("Enter 11 for EXERCISE 4.6 Delete from String Map: ")
	fmt.Println("Enter 12 for EXERCISE 5.1 basic calculator: ")
	fmt.Println("Enter 13 for EXERCISE 5.2 fibinocci series: ")
	fmt.Println("Enter 14 for EXERCISE 5.3 Own function parameters that consisting of an anonymous function: ")
	fmt.Println("Enter 15 for EXERCISE 6.1 Employee struct: ")
	fmt.Println("Enter 16 for EXERCISE 6.2 Area and Perimeter of Rectangle: ")
	fmt.Println("Enter 17 for EXERCISE 6.3 Area and Perimeter of Square: ")
	fmt.Println("Enter 18 for EXERCISE 6.4 Area and Perimeter of Circle: ")
	fmt.Println("Enter 19 for EXERCISE 6.5 Hello World using non struct Method: ")
	fmt.Println("Enter 20 for EXERCISE 6.6 Reverse a string non struct Method.: ")
	fmt.Scanf("%d", &c)

	switch c {
	case 1:
		// Exercise 1.1 Hello world
		fmt.Println(basic.Hello())
	case 2:
		// Exercise 3.1 : Modulo operation
		func(x int) {
			for i := 10; i <= 100; i++ {
				remainder := i % x
				fmt.Printf("The remainder for %d when divided by %d is %d \n", i, x, remainder)
			}
		}(4)
	case 3:
		// Exercise 3.2 : Modulo operation
		func(start, end, x, y int) {
			fmt.Printf("The numbers which are divisible by %d and multiple of %d, between %d and %d  are \n", x, y, start, end)
			for i := start; i <= end; i++ {
				if i%x == 0 && i%y == 0 {
					fmt.Println(i)
				}
			}
		}(1500, 2700, 7, 5)
	case 4:
		//EXERCISE 4.1: Multidimensional Slice
		basic.MultiDSlice()
	case 5:
		//EXERCISE 4.2: Slice of int
		a, b, c := basic.SliceInt()
		fmt.Println(a, "\n", b, "\n", c)
	case 6:
		//EXERCISE 4.3:Delete Slice of int
		z := basic.SliceIntDel()
		fmt.Println(z)
	case 7:
		//EXERCISE 4.4 Dynamic slice & convert slice to Map
		basic.DynamicSlice()
	case 8:
		//EXERCISE 4.4 Create a menu with add,delete option for slice & map
		basic.SliceMenu()
	case 9:
		//EXERCISE 4.5 String Map
		basic.MapStr()
	case 10:
		//EXERCISE 4.6 Add to String Map
		basic.AddToMapStr()

	case 11:
		//EXERCISE 4.6 Delete from String Map
		basic.DltFrmMapStr()
	case 12:
		//EXERCISE 5.1 basic calculator
		basic.Calculator()
	case 13:
		//EXERCISE 5.2 fibinocci series
		basic.Fib(20)
	case 14:
		//EXERCISE 5.3 program having your own function parameters that consisting of an anonymous function.
		basic.OwnAnonymFunc()
	case 15:
		//EXERCISE 6.1 struct EMPLOYEE.
		basic.NewEmployee("Bob", "GO", 1225)
	case 16:
		//EXERCISE 6.2 Area and Perimeter of Rectangle.
		basic.NewRect(10.0, 5.0)
	case 17:
		//EXERCISE 6.3 Area and Perimeter of Square.
		basic.NewSquare(3.0)
	case 18:
		//EXERCISE 6.4 Area and Perimeter of Circle.
		basic.NewCircle(5.0)
	case 19:
		//EXERCISE 6.5 non struct Method.
		basic.PrintText()
	case 20:
		//EXERCISE 6.6 Reverse a string non struct Method.
		basic.PrintRevText()
	default:
		fmt.Println("Invalid Choice")
	}
}

func exMedium() {
	var c int

	fmt.Println("Choose Medium Exercises: ")
	fmt.Println("Enter 1 for Exercise 7.1 Pointers Basic: ")
	fmt.Println("Enter 2 for Exercise 7.2 : Average mark using a pointer receiver: ")
	fmt.Println("Enter 3 for Exercise 8.1 : Print Area of shape passed: ")
	fmt.Println("Enter 4 for Exercise 8.2 : Print Area & Perimeter of shape passed: ")
	fmt.Println("Enter 5 for Exercise 8.3 : Print Total Area of All shapes passed: ")
	fmt.Println("Enter 6 for Exercise 8.4 : Print Total Expense for a Company: ")
	fmt.Println("Enter 7 for Exercise 8.5 : Empty Interface: ")
	fmt.Println("Enter 8 for Exercise 9.1 : Write sorted input strings to a file: ")
	fmt.Println("Enter 9 for Exercise 9.2 : Convert struct to JSON and write to a Json file: ")
	fmt.Println("Enter 10 for Exercise 9.3 : Output either json/xml based on conf choice: ")
	fmt.Scanf("%d", &c)

	switch c {
	case 1:
		// Exercise 7.1 Pointers Basic
		medium.PntrBasic()
	case 2:
		// Exercise 7.2 Average mark using a pointer receiver.
		medium.MakeStdnts()
	case 3:
		// Exercise 8.1 : Print Area of shape passed.
		medium.PrntAreaOfShape()
	case 4:
		// Exercise 8.2 Print Area & Perimeter of All shapes passed
		medium.PrntAreaPermtrOfShape()
	case 5:
		// Exercise 8.3 Print Total Area of All shapes passed
		medium.PrntTotalAreaOfShapes()
	case 6:
		// Exercise 8.4 Print Total Expense for a Company
		medium.PrntTotalExpns()
	case 7:
		// Exercise 8.5 : Empty Interface.
		medium.EmptyIntrfc()
	case 8:
		// Exercise 9.1 : Write sorted input strings to a file.
		medium.WrtToFile()
	case 9:
		// Exercise 9.2 : Write struct to a json file.
		medium.WrtToJsonFile()
	case 10:
		// Exercise 9.3 : Output either json/xml based on conf choice
		medium.WrtToDynFile()
	default:
		fmt.Println("Invalid Choice")
	}
}

func menu() {
	var c int
	fmt.Println("Select Difficulty Level: ")
	fmt.Println("Enter 1 for Basics: ")
	fmt.Println("Enter 2 for Normal: ")
	fmt.Println("Enter 3 for Advanced: ")
	// fmt.Println("Enter 4 for Challenging: ")
	fmt.Scanf("%d", &c)

	switch c {
	case 1:
		exBasic()
	case 2:
		exMedium()
	case 3:
		advanced.Menu()
	default:
		fmt.Println("Invalid Choice")
	}
}
