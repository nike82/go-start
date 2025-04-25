package basics

import "fmt"

func main() {

	mess := "hello world!"

	for i, v := range mess {
		//fmt.Println(i, v)
		fmt.Printf("index: %d, rune: %c\n", i, v)
	}
}
