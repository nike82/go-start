package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "hello \ngo!"
	message1 := "hello \tgo!"
	message2 := "hello \rgo!"
	rawMessage := `hello \ngo!`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("lengths of message is", len(message))
	fmt.Println("lengths of message is", len(message1))
	fmt.Println("lengths of message is", message2[0])

	greeting := "hello"
	name := "nike"
	fmt.Println(greeting + name)
	fmt.Println(greeting < name)

	for i, char := range message {
		fmt.Printf("Char at index %d is %c\n", i, char)

	}

	for _, char := range message {
		fmt.Printf("%v\n", char)
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("type of cstr is %T\n", cstr)

	kolia:="コリャ"
	fmt.Println(kolia)

	for _, runeValue:= range kolia{
		fmt.Printf("%c\n", runeValue)
	}
}
