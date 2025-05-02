package intermediate

import "fmt"

func main() {

	num := 42345345
	fmt.Printf("%05d\n", num)

	message:= "hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1:= "Hello \nWorld!"
	message2:= `Hello \nWorld!`
	fmt.Println(message1)
	fmt.Println(message2)

	sqlquery:= `SELECT * FROM users WHERE age > 30`
	fmt.Println(sqlquery)
}
