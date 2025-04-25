package basics

import "fmt"

func main() {
	fmt.Println(add(1, 2))

	hello := func() {
		fmt.Println("hello")
	}

	hello()

	operation := add
	result := operation(3, 4)
	fmt.Println(result)

	//passing func as an arg
	res := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", res)

	//return and use func

	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 x 2 = ", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// func that takes a func as an arg
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// func that returns a func

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
