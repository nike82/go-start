package intermediate

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	subtraction := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtraction(1))
	fmt.Println(subtraction(2))
	fmt.Println(subtraction(3))
	fmt.Println(subtraction(4))
	fmt.Println(subtraction(5))
}

func adder() func() int {
	i := 0
	fmt.Println("previouse value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
