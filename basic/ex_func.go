package basic

import "fmt"

func add(a, b int) (int, string) {
	return a + b, "+"
}

func sub(a, b int) (int, string) {
	return a - b, "-"
}

func mul(a, b int) (int, string) {
	return a * b, "*"
}

func div(a, b int) (int, string) {
	return a / b, "/"
}

func process(f func(int, int) (int, string), a, b int) {
	res, operator := f(a, b)
	fmt.Printf("%d %s %d = %d \n", a, operator, b, res)
}

func Calculator() {
	var a, b, c int

	fmt.Println("Choose any one operation you wish to perform: ")
	fmt.Println("Enter 1 for Addition: ")
	fmt.Println("Enter 2 for Subtraction: ")
	fmt.Println("Enter 3 for Multiplication: ")
	fmt.Println("Enter 4 for Division: ")
	fmt.Scanf("%d", &c)

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &b)

	switch c {
	case 1:
		process(add, a, b)
	case 2:
		process(sub, a, b)
	case 3:
		process(mul, a, b)
	case 4:
		process(div, a, b)
	default:
		fmt.Println("Invalid Choice")
	}
}

func Fib(n int) {
	t1 := 0
	t2 := 1
	fmt.Printf("First %d terms: ", n)
	for i := 0; i < n; i++ {
		fmt.Print(t1, " , ")
		sum := t1 + t2
		t1 = t2
		t2 = sum
	}
	print("\n")
}

func greatNewbie(i func(p, q string) string) {
	fmt.Println(i("Welcome folks ", "to the "))

}

func OwnAnonymFunc() {
	value := func(p, q string) string {
		return p + q + "Go World!!"
	}
	greatNewbie(value)
}

func Hello() string {
	return "Hello World"
}

func Prod(m map[int]int) int {
	prod := 1
	for _, val := range m {
		prod = prod * val
	}
	fmt.Printf("The prod is :%d \n", prod)
	return prod
}
