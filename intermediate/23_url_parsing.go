package intermediate

import (
	"fmt"
	"net/url"
)

func main() {

	rawURL := "https://example.com:8080/path?query=param#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=Nike&age=43"
	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "Nike")
	query.Set("age", "34")
	query.Set("ha", "ha")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Build URL:", baseURL.String())

	values := url.Values{}
	values.Add("name", "Jimm")
	values.Add("age", "22")
	values.Add("city", "batum")
	values.Add("contry", "GE")
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)
	baseURL1 := "https://myurl.com/blah"
	fullURL := baseURL1 + "?" + encodedQuery
	fmt.Println(fullURL)
}
