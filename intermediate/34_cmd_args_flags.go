package intermediate

import (
	"flag"
	"fmt"
	//"os"
)

func main() {

	//fmt.Println("Command:", os.Args[0])
	// for i, arg := range os.Args {
	// 	fmt.Println("Args:", i, ":", arg)
	// }
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 27, "Age of user")
	flag.BoolVar(&male, "male", true, "Gender of the user")
	flag.Parse()
	fmt.Println("Name", name)
	fmt.Println("Age", age)
	fmt.Println("Male", male)
}
