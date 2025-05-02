package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person //anonymouse name
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("I'm %s and I'm %d years old\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("I'm %s employee id : %s and my salary %.2f\n", e.name, e.empId, e.salary)
}

func main() {
	emp := Employee{
		person: person{
			name: "John",
			age:  30,
		},
		empId:  "E001",
		salary: 5000,
	}
	fmt.Println("name:", emp.name)
	fmt.Println("age:", emp.age)
	fmt.Println("id:", emp.empId)
	fmt.Println("salary:", emp.salary)

	emp.introduce()
}
 