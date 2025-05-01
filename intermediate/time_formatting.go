package intermediate

import (
	"fmt"
	"time"
)

func main() {

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2023-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}
	fmt.Println(t)

	str1 := "Aug 05, 2022 04:12 PM"
	layout1 := "Jan 02, 2006 03:04 PM"
	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)
}
