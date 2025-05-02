package intermediate

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	fmt.Println(someOfDigits(9))
	fmt.Println(someOfDigits(12))
	fmt.Println(someOfDigits(12345))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func someOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + someOfDigits(n/10)
}
