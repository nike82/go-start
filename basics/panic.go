package basics

import "fmt"

func main() {

	//panic(interface{})
	process(1)
	process(-1)

}


func process(input int) {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")

	if input < 0 {
		fmt.Println("before panic")
		panic("input must be > 0")
		//fmt.Println("after panic")
		//defer fmt.Println("defer3")
	}
	fmt.Println("processing input:", input)
}
