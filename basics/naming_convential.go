package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	const MAXTRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
