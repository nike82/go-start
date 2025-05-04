package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")
	fmt.Println(user)
	fmt.Println(home)

	err := os.Setenv("ME", "NIKE")
	if err != nil {
		fmt.Println("Error setting env var:", err)
	}
	fmt.Println(os.Getenv("ME"))
	for _, e := range os.Environ() {
		kvPair := strings.SplitN(e, "=", 2)
		fmt.Println(kvPair[0])
	}
	err = os.Unsetenv("ME")
	if err != nil {
		fmt.Println("Error unsetting env var:", err)
		return
	}
	fmt.Println(os.Getenv("ME"))
	str:="a=b=c=d=f"
	fmt.Println(strings.SplitN(str,"=", -1))
	fmt.Println(strings.SplitN(str,"=", 0))
	fmt.Println(strings.SplitN(str,"=", 1))
	fmt.Println(strings.SplitN(str,"=", 2))
	fmt.Println(strings.SplitN(str,"=", 3))
	fmt.Println(strings.SplitN(str,"=", 4))
}
