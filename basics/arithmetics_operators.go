package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Substruction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	// signed
	var maxInt int64 = 9223372036854775807 // max int64
	fmt.Println(maxInt)

	maxInt++
	fmt.Println(maxInt)

	// unsigned
	var uMaxInt uint64 = 18446744073709551615 // max uint64
	fmt.Println(uMaxInt)

	uMaxInt++
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
