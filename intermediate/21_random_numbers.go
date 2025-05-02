package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// fmt.Println(rand.Intn(101))
	// fmt.Println(rand.Intn(101))
	// fmt.Println(rand.Intn(101))

	val := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(val.Intn(6) + 5)
	fmt.Println(val.Intn(6) + 5)
	fmt.Println(val.Intn(6) + 5)

	fmt.Println(rand.Intn(6) + 5)
	fmt.Println(rand.Float64()) //[0.0, 1.0)

	for {
		fmt.Println("Welcome to the dice game!")
		fmt.Println("1. roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2):")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid!")
			continue
		}
		if choice == 2 {
			fmt.Println("Goodbye!")
			break
		}

		diel := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("You rolled %d and a %d.\n", diel, die2)
		fmt.Println("Total:", diel+die2)
		fmt.Print("Do you want again?")

		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("invalid")
			break
		}
		if rollAgain =="n" {
			fmt.Println("Bye!")
			break
		}
	}

}
