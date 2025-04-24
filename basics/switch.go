package basics

import "fmt"

func main() {
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("apple")
	case "banana":
		fmt.Println("banana")
	default:
		fmt.Println("no clue")
	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Thursday", "Friday":
		fmt.Println("weekday")
	case "Sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("no clue")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("less then 10")
	case number >= 10 && number < 20:
		fmt.Println("between 10 and 20")
	default:
		fmt.Println("more 20")
	}

	number2 := 2

	switch {
	case number2 > 1:
		fmt.Println("more 1")
		fallthrough
	case number2 == 2:
		fmt.Println("equal 2")
	default:
		fmt.Println("not 2")
	}

	checkType(10)
	checkType(1.2)
	checkType("eee")
	checkType(false)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
		// can't fallthrough
	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("no clue")
	}
}
