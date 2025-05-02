package intermediate

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix time: ", unixTime)

	t:= time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time: ", t.Format("2006/01/02"))
}
