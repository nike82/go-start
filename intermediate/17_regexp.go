package intermediate

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("he sad, \"I am great\"")
	fmt.Println(`he sad, "I am great"`)
	fmt.Println("he sad, \n\"I am great\"")

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_email.com"

	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2025-04-29"

	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	str:= "Hello World!"
	re = regexp.MustCompile(`[aeiou]`)
	result:=re.ReplaceAllString(str, "*")
	fmt.Println(result)

	re = regexp.MustCompile(`(?i)go`)
	text:= "Golang is great"
	fmt.Println("match:", re.MatchString(text))
}
