package advance

import "fmt"

func main() {
	ch := make(chan int)
	produser(ch)
	consumer(ch)
}

func produser(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("Received: ", v)
	}
}
