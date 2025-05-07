package advance

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 sec")
	}()

	go func() {
		//ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 sec")
	}()
	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println("end")

}
