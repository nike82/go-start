package intermediate

import "fmt"

func main() {

	fmt.Print("hello")
	fmt.Print("nike")
	fmt.Print(35, 789)

	fmt.Println("hello")
	fmt.Println("nike")
	fmt.Println(35, 789)

	s := fmt.Sprint("hello", "nike", 12)
	fmt.Println(s)

	fmt.Println("hello", "nike", 12)

	var name string
	var age int

	fmt.Print("enter name and age:")
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s, age: %d\n", name, age)

	fmt.Print("enter name and age:")
	fmt.Scanln(&name, &age)
	fmt.Printf("Name: %s, age: %d\n", name, age)

	err := chaeckAge(15)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func chaeckAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young", age)
	}
	return nil
}
