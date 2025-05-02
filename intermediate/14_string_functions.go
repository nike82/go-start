package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "Go!"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])
	fmt.Println(str[1:5])

	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"Germany", "France", "Italy"}

	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Contains(str, "Go?"))

	replaced := strings.Replace(str, "Go", "Nike", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone! "
	fmt.Println(strings.TrimSpace(strwspace))
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo ", 4))
	fmt.Println(strings.Count("foo", "o"))
	fmt.Println(strings.HasPrefix("foo", "fo"))
	fmt.Println(strings.HasSuffix("foo", "oo"))

	str5 := "Hel1lo, 123 Go 33!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello デイズ"
	fmt.Println(utf8.RuneCountInString(str6))

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	result = builder.String()
	fmt.Println(result)

	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	builder.Reset()
	builder.WriteString("Starting a new string")
	result = builder.String()
	fmt.Println(result)
}
