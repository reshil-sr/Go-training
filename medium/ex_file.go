package medium

import (
	"encoding/json"
	"fmt"
	"gotraining/basic"
	"io/ioutil"
	"os"
	"sort"
)

type person struct {
	Name string
	Age  int
}

//Exercise 9.1
func WrtToFile() {
	fmt.Println("Write inputed strings to a file after sorting")
	dynamicSlice()
}

func dynamicSlice() {
	fmt.Println("Enter Strings:")
	inputSlice := basic.UsrInput([]string{}, nil)
	fmt.Println("Input Slice:", inputSlice)
	sort.Strings(inputSlice)
	fmt.Println("Sorted Slice:", inputSlice)
	file, err := os.Create("./files/input.txt")

	chkNilErr(err)

	defer file.Close()
	for k, value := range inputSlice {
		strKey := fmt.Sprintf("%d", k)

		// fmt.Fprintln(file, strKey+")"+value)
		_, err := file.WriteString(strKey + ")" + value + "\n")

		chkNilErr(err)
	}

}

//Exercise 9.2
func WrtToJsonFile() {
	fmt.Println("Convert struct to JSON and write to a Json file")
	procsJson()
}

func procsJson() {
	prsn_1 := &person{"John", 30}
	e, err := json.Marshal(prsn_1)
	chkNilErr(err)
	fmt.Println(string(e))
	file, _ := json.Marshal(string(e))
	_ = ioutil.WriteFile("./files/input.json", file, os.ModePerm)
}

func chkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

//Exercise 9.3
func WrtToDynFile() {
	fmt.Println("Output either json/xml based on conf choice")
}
