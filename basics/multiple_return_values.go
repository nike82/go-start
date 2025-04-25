package basics

import (
	"errors"
	"fmt"
)

func main() {
	q, r := divide(10, 3)
	fmt.Printf("quest: %d, reminder: %d\n", q, r)

	res, err := compare(3, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}

func divide(a, b int) (quotient int, reminder int) {
	quotient = a / b
	reminder = a % b
	return  
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a > b", nil
	} else if b > a {
		return "b > a", nil
	} else {
		return "", errors.New("Can't compare")
	}
}
