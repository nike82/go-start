package advance

import (
	"fmt"
	"time"
)

func main() {
	//##############################
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// //ch <- 3

	// fmt.Println("Value: ", <-ch)
	// fmt.Println("Value: ", <-ch)
	// ch <- 3
	// fmt.Println("Buffered channels")
	//##############################
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		//fmt.Println("Goroutine 2sec")
		time.Sleep(1 * time.Second)
		fmt.Println("Received:", <-ch)
		//ch <- 3
	}()
	//fmt.Println("Blocking starts")
	ch <- 3
	//fmt.Println("Blocking ends")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Buffered channels")
	//##############################

	// ch := make(chan int, 2)

	// go func ()  {
	// 	time.Sleep(2*time.Second)
	// 	ch <- 1
	// 	ch <- 2
	// }()
	// fmt.Println("Value: ", <-ch)
	// fmt.Println("Value: ", <-ch)
	// fmt.Println("end")
}
