package basics

import (
	"fmt"
	"maps"
)

func main() {

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["k"])
	myMap["key2"] = 100
	fmt.Println(myMap["key2"])

	delete(myMap, "key1")

	fmt.Println(myMap)

	myMap["key3"] = 19
	myMap["key4"] = 91
	myMap["key5"] = 92

	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println(ok)
	} else {
		fmt.Println("nope")
	}
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("equal")
	}

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("no nil")
	}

	val:= myMap4["key"]
	fmt.Println(val)

	//val:= myMap4["key"] = "val"

	myMap4= make(map[string]string)
	myMap4["key"] = "val"
	fmt.Println(myMap4)

	fmt.Println(len(myMap4))

	myMap5:= make(map[string]map[string]string)

	myMap5["map1"] = myMap4

	fmt.Println(myMap5)
}
