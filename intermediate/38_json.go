package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{FirstName: "Nike", Age: 37, EmailAddress: "test@mail.com"}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		FirstName:    "John",
		Age:          23,
		EmailAddress: "john@mail.com",
		Address: Address{
			City:  "BA",
			State: "ADJ",
		},
	}
	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	fmt.Println(string(jsonData1))

	jsonData2 := `{"full_name": "Jenny Doe", "emp_id": "0909", "age": 30, "address": {"city": "NY", "state": "NY"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Age:", employeeFromJson.Age)
	fmt.Println("City:", employeeFromJson.Address.City)

	listOfCityState:=[]Address{
		{City: "NY", State: "NY"},
		{City: "San Jose", State: "CA"},
	}
	fmt.Println(listOfCityState)
	jsonList, err:= json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error marshalling to JSON:", err)
	}
	fmt.Println("JSON list", string(jsonList))

	jsonData3:= `{"name": "John", "age": 30, "address": {"city": "NY", "state": "NY"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling JSON:", err)
	}
	fmt.Println("Unmurshaled JSON ", data)
	fmt.Println("Unmurshaled JSON ", data["address"])
	fmt.Println("Unmurshaled JSON ", data["name"])
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpId    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
