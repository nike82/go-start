package basics

import "fmt"

func main() {
	statement, total := sum("The sum of 1,2,3 is", 1, 2, 3)
	fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5}

	sequence, total2:= sum2(3, numbers...)
	fmt.Println("Sequence:", sequence, "total:", total2)

}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}

func sum2(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
