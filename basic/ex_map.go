package basic

import "fmt"

var strMap = map[string]string{
	"Griesemer_Robert": "V8 JavaScript engine",
	"Pike_Rob":         "The Unix Programming Environment",
	"Thompson_Ken":     "unix os",
}

func MapStr() {
	i := 0
	for a, b := range strMap {
		fmt.Printf("%d) %v:%v \n", i, a, b)
		i++
	}
}

func AddToMapStr() {
	fmt.Println("Number of Enteries: ")
	var lmt int
	fmt.Scanf("%d", &lmt)
	for i := 0; i < lmt; i++ {
		fmt.Println("Enter last_first Name : likes")
		var lstFrstName string
		var likes string
		fmt.Scanf("%s : %s", &lstFrstName, &likes)
		strMap[lstFrstName] = likes
	}
	MapStr()
}

func DltFrmMapStr() {

	fmt.Println("Enter key to deletefrom Map")
	var key string
	fmt.Scanf("%s", &key)
	fmt.Printf("Deleting the key named %s from the map \n", key)
	delete(strMap, key)
	MapStr()
}
