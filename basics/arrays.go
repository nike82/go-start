package basics

import "fmt"

func main() {

	var nums [5]int
	fmt.Println(nums)

	nums[4] = 20
	fmt.Println(nums)

	fruits := [4]string{"apple", "banana", "orange", "peach"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("3:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("orig:", originalArray)
	fmt.Println("copy:", copiedArray)

	var pointerToArray *[3]int
	pointerToArray = &originalArray

	fmt.Println("orig:", originalArray)
	fmt.Println("copy:", *pointerToArray)

	for i := 0; i < len(nums); i++ {
		fmt.Println("element at index:", i, ":", nums[i])
	}

	for i, v := range nums {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	for _, v := range nums {
		fmt.Printf("value: %d\n", v)
	}

	a, _ := someFunc()
	fmt.Println(a)

	b := 2
	_ = b

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1 == arr2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)
}

func someFunc() (int, int) {
	return 1, 2
}
