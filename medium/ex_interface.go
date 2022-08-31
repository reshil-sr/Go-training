package medium

import "fmt"

const Pi = 3.142

type shape interface {
	area() float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

// rectangle
type rectangle struct {
	length, breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

// square
type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) perimeter() float64 {
	return 4 * s.length
}

// circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * Pi * c.radius
}

// misc
type misc struct {
	shapes   []shape
	bgcolour string
	fgcolour string
}

func (m misc) totalArea() {
	area := 0.0
	for _, shp := range m.shapes {
		area += shp.area()
	}
	fmt.Printf("Total Area of all Shapes is %.2f \n", area)
}

type salary interface {
	totalSalary() float32
}

type employee struct {
	empld    int
	basicPay float32
}

func (e employee) totalSalary() float32 {
	return e.basicPay
}

type permEmployee struct {
	e  employee
	pf float32
}

func (pe permEmployee) totalSalary() float32 {
	return pe.e.basicPay + pe.pf
}

//Exercise 8.1
func info(shp shape) {
	fmt.Printf("Area of %T is %v \n", shp, shp.area())
}

func PrntAreaOfShape() {
	fmt.Printf("Exercise 8.1 : Print Area of shape passed. \n")
	c := circle{radius: 3}
	s := square{length: 5}
	rect := rectangle{length: 3, breadth: 5}
	info(c)
	info(s)
	info(rect)
}

//Exercise 8.2
func areaPermtrOfShape(g geometry) {
	fmt.Printf("Area of %T is %v \n", g, g.area())
	fmt.Printf("Perimeter of %T is %v \n", g, g.perimeter())
}

func PrntAreaPermtrOfShape() {
	fmt.Printf("Exercise 8.2 : Print Area & Perimeter of shape passed. \n")
	c := circle{radius: 3}
	s := square{length: 5}
	rect := rectangle{length: 3, breadth: 5}
	areaPermtrOfShape(c)
	areaPermtrOfShape(s)
	areaPermtrOfShape(rect)
}

//Exercise 8.3
func PrntTotalAreaOfShapes() {
	fmt.Printf("Exercise 8.3 Print Total Area of All shapes passed. \n")
	msc := new(misc)
	c := circle{radius: 3}
	s := square{length: 5}
	rect := rectangle{length: 3, breadth: 5}
	msc.shapes = []shape{c, s, rect}
	msc.bgcolour = "red"
	msc.fgcolour = "black"
	msc.totalArea()
}

//Exercise 8.4
func PrntTotalExpns() {
	fmt.Printf("Exercise 8.4 Calculate Total Expense for a Company. \n")
	cntrE1 := new(employee)
	cntrE1.empld = 201
	cntrE1.basicPay = 20000

	cntrE2 := new(employee)
	cntrE2.empld = 202
	cntrE2.basicPay = 25500

	prmE1 := new(permEmployee)
	prmE1.e = employee{101, 30000}
	prmE1.pf = 1800

	prmE2 := new(permEmployee)
	prmE2.e = employee{102, 32000}
	prmE2.pf = 1800

	prmE3 := new(permEmployee)
	prmE3.e = employee{103, 33200}
	prmE3.pf = 1800

	result := calcCTC(cntrE1, cntrE2, prmE1, prmE2, prmE3)
	fmt.Printf("Total cost to company is: %.2f \n", result)

}

func calcCTC(sals ...salary) float32 {
	var tlCTC float32
	tlCTC = 0
	for _, sal := range sals {
		tlCTC += sal.totalSalary()
	}
	return tlCTC
}

// Exercise 8.5
func prntType(i interface{}) {
	fmt.Printf("%#v %T\n", i, i)
}

func EmptyIntrfc() {
	fmt.Printf("Print value and type of 3 different types of inputs \n")
	types := []interface{}{1, "two", 3.5}
	for _, val := range types {
		prntType(val)
	}
}
