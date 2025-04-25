package basics

import "fmt"


func main() {

	process(2)
	
}

func process(i int){
	defer fmt.Println("defer value", i)
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	i++
	fmt.Println("normal exec state")
	fmt.Println("i value", i)
}

