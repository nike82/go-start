package main

import "fmt"

func main() {

	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	temperature := 25
	if temperature >= 30 {
		fmt.Println("Hot!")
	} else {
		fmt.Println("Cool!")
	}

	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	if score >= 90 {
		fmt.Println("A")
	}
	if score >= 80 {
		fmt.Println("B")
	}
	if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	num := 15

	if num % 2 == 0 {
		if num % 3 == 0 {
			fmt.Println("Number is divisible by 2 and 3")
		} else {
			fmt.Println("Number is divisible only by 2")
		}
	} else {
		fmt.Println("Number is not divisible only by 2")
	}

	if 10%2 ==0 || 5%2 ==0 {
		fmt.Println("Either 10 or 5 are even")
	}

	if 10%2 ==0 && 6%2 ==0 {
		fmt.Println("Both 10 or 6 are even")
	}
}
