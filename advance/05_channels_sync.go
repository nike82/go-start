package advance

import (
	"fmt"
	"time"
)

func main() {

	// done:=make(chan struct{})
	// go func ()  {
	// 	fmt.Println("Start...")
	// 	time.Sleep(2*time.Second)
	// 	done <- struct{}{}
	// }()
	// <- done
	// fmt.Println("Finished!")

	// ch:=make(chan int)
	// go func ()  {
	// 	ch <- 9
	// 	time.Sleep(1*time.Second)
	// 	fmt.Println("Sent value")
	// }()
	// value:= <- ch
	// fmt.Println(value)
	// fmt.Println("Finished!")

	// numGoroutines := 3
	// done := make(chan int, 3)
	// for i := range numGoroutines {
	// 	go func(id int) {
	// 		fmt.Printf("Goroutin %d working... \n", id)
	// 		time.Sleep(time.Second)
	// 		done <- id
	// 	}(i)
	// }
	// for range 3 {
	// 	<- done
	// }
	// fmt.Println("Finished!")

	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- "Hello " + string('0'+i)
			time.Sleep(100 * time.Microsecond)
		}
		close(data)
	}()
	
	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}
	fmt.Println("Finished!")
}
