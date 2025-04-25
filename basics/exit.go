package basics

import (
	"fmt"
	"os"
)


func main() {

	defer fmt.Println("defer")
	fmt.Println("start the main func")
	os.Exit(1)
	fmt.Println("never exec")
}