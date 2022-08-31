package medium

import "fmt"

var pntrVar = 10

type student struct {
	student_id   int
	student_name string
	subject_name string
	mark         []int
}

func changeVal(i *int) {
	*i = 100
}

func PntrBasic() {
	fmt.Printf("Initital value of variable is %d \n", pntrVar)
	fmt.Println("The pointer is : ", &pntrVar)
	changeVal(&pntrVar)
	fmt.Printf("After update using pointer the value of variable is %d \n", pntrVar)
}

func MakeStdnts() {
	std1 := new(student)
	std1.student_id = 1
	std1.student_name = "Alen"
	std1.subject_name = "GO"
	std1.mark = []int{50, 60, 55, 75}
	std2 := student{2, "Bob", "Python", []int{30, 40, 50, 60}}
	std3 := student{2, "Charle", "C++", []int{33, 47, 52, 35}}

	fmt.Printf("The average mark of %s is %.2f \n", std1.student_name, std1.avgMark())
	fmt.Printf("The average mark of %s is %.2f \n", std2.student_name, std2.avgMark())
	fmt.Printf("The average mark of %s is %.2f \n", std3.student_name, std3.avgMark())

}

func (std student) avgMark() float32 {
	sum := 0
	for _, s := range std.mark {
		sum += s
	}
	avg := float32(sum) / float32(len(std.mark))
	return avg
}
