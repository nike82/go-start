package advance

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received message:", msg)
	// default:
	// 	fmt.Println("No message available")
	// }

	// select {
	// case ch <- 1:
	// 	fmt.Println("Send message")
	// default:
	// 	fmt.Println("Channel is not ready to receive")
	// }

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping")
				return
			default:
				fmt.Println("Waiting for data")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
