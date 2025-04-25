package basics

import "fmt"

func main() {
	process()
	fmt.Println("return from process")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()

	fmt.Println("start process")
	panic("smth wrong")
	fmt.Println("end process")

}
