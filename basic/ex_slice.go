package basic

import "fmt"

var intSlice = []int{1, 3, 5, 7, 9, 10, 12, 14, 16, 18}
var inputSlice = []string{}
var strMapFrmSlice = make(map[string]string)

func MultiDSlice() {

	twoDSlice := make([][]string, 2)
	twoDSlice[0] = []string{"James", "Bond", "Shaken, not stirred"}
	twoDSlice[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	for _, row := range twoDSlice {
		for _, val := range row {
			fmt.Printf(val + "\t|")
		}
		fmt.Printf("\n")
	}
}

func SliceInt() ([]int, []int, []int) {

	for val := range intSlice {
		fmt.Printf("%d is of type %T and its value is of type %T\n", val, intSlice, val)
	}

	sLen := len(intSlice)

	return intSlice[:sLen-1], intSlice[3:sLen], intSlice[2 : sLen-2]
}

func SliceIntDel() []int {

	intSlice = append(intSlice[:2], intSlice[6:]...)

	return intSlice
}

func UsrInput(inputStr []string, err error) []string {
	if err != nil {
		return inputStr
	}
	var val string
	n, err := fmt.Scanf("%s", &val)
	if n == 1 {
		inputStr = append(inputStr, val)
	}
	return UsrInput(inputStr, err)
}

func DynamicSlice() {
	fmt.Println("Enter Strings:")
	inputSlice = UsrInput([]string{}, nil)
	fmt.Println("Input Slice:", inputSlice)
	fmt.Println("The input slice is converted to Map:")
	sliceToMap(inputSlice)
}

func sliceToMap(input []string) {
	for key, s := range input {
		strKey := fmt.Sprintf("%d", key)
		strMapFrmSlice[strKey] = s
	}
	fmt.Println(strMapFrmSlice)
}

func dynSliceAppend() {
	fmt.Println("Enter Strings:")
	inputSlice = UsrInput(inputSlice, nil)
	fmt.Println("Input Slice:", inputSlice)
	fmt.Println("The input slice is converted to Map:")
	sliceToMap(inputSlice)
}

func dynMapAppend() {
	var key string
	var value string
	fmt.Println("Enter Map key")
	fmt.Scanf("%s", &key)
	fmt.Println("Enter Map value")
	fmt.Scanf("%s", &value)
	strMapFrmSlice[key] = value
	fmt.Println(strMapFrmSlice)
}

func dynSliceDel(s []string) {
	var key int
	fmt.Println("Enter key to Delete")
	fmt.Scanf("%d", &key)
	s[key] = s[len(s)-1]
	inputSlice = s[:len(s)-1]
	fmt.Println(inputSlice)
}

func dynMapDelete() {
	fmt.Println("Enter key to deletefrom Map")
	var key string
	fmt.Scanf("%s", &key)
	fmt.Printf("Deleting the key named %s from the map \n", key)
	delete(strMapFrmSlice, key)
	fmt.Println(strMapFrmSlice)
}

func SliceMenu() {

	var c int

	fmt.Println("Choose any one operation you wish to perform: ")
	fmt.Println("Enter 1 for Append value to slice: ")
	fmt.Println("Enter 2 for Append a key value pair to map: ")
	fmt.Println("Enter 3 for Delete given value from slice: ")
	fmt.Println("Enter 4 for Delete value from map, with the given key.: ")
	fmt.Scanf("%d", &c)

	switch c {
	case 1:
		dynSliceAppend()
	case 2:
		dynMapAppend()
	case 3:
		dynSliceDel(inputSlice)
	case 4:
		dynMapDelete()
	default:
		fmt.Println("Invalid Choice")
	}
}
