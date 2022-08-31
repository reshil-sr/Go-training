package basic

import "fmt"

const Pi = 3.142

type employee struct {
	name   string
	code   int
	domain string
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, breadth float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type text string

func (r rectangle) area() float64 {
	return r.length * r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) perimeter() float64 {
	return 4 * s.length
}

func (c circle) area() float64 {
	return Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * Pi * c.radius
}

func NewEmployee(name, domain string, code int) {
	fmt.Println(employee{name: name, code: code, domain: domain})
}

func NewRect(len, brdth float64) {
	rect := rectangle{length: len, breadth: brdth}
	fmt.Println("Area of rectangle is ", rect.area())
	fmt.Println("Perimeter of rectangle is ", rect.perimeter())
}

func NewSquare(len float64) {
	s := square{length: len}
	fmt.Println("Area of Square is ", s.area())
	fmt.Println("Perimeter of Square is ", s.perimeter())
}

func NewCircle(radius float64) {
	c := circle{radius: radius}
	fmt.Println("Area of circle is ", c.area())
	fmt.Println("Perimeter of circle is ", c.perimeter())
}

func (txt text) concatText(txt2 text) {
	res := txt + " " + txt2
	fmt.Printf("%s is of type %T \n", res, res)
}

func PrintText() {
	t := text("Hello")
	t.concatText(text("world"))
}

func (txt text) reverse() (result text) {
	for _, v := range txt {
		result = text(v) + result
	}
	return result
}

func PrintRevText() {
	t := text("DLROW OLLEH")
	revTxt := t.reverse()
	fmt.Printf("The reverse of %s is %s which is of type %T \n", t, revTxt, revTxt)
}
