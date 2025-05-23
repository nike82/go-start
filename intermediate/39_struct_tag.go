package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	Age       int    `json:"-"`
}

func main() {
	//person := Person{FirstName: "John", LastName: "Full", Age: 43}
	person := Person{FirstName: "John", Age: 30}
	jsonData, err:= json.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling struct", err)
	}
	fmt.Println(string(jsonData))
}
