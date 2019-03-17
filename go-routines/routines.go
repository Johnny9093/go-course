package go_routines

import (
	"fmt"
	"time"
)

func main() {
	// use channels to communicate between different go routines
	cn := make(chan int, 10)
	// cn <- 2 // push value to channel

	// 10 workers that will print their "id" and "task id".
	// Subscriber
	for i := 0; i < 10; i++ {
		go func(index int, queue chan int) {
			for {
				fmt.Printf("Worker %d: %d\n", index, <- queue)
			}
		}(i, cn)
	}

	// Publisher
	for i := 0; i < 20; i++ {
		cn <- i // will be blocked if no one empties the channel?
	}

	//// run for as go routine
	//go func(inputChannel chan int) {
	//	println("starting goroutine")
	//	for {
	//		println("pulling value")
	//		value := <- inputChannel // since there are no values in the channel, the func is blocked here
	//		// <- chan (pulling from chanel) is blocking.
	//		println(value)
	//	}
	//}(cn)
	//// exits immediately, we wont see the result

	//time.Sleep(3 * time.Second)
	//cn <- 13

	time.Sleep(1 * time.Second)

	// main routine ended, all goroutines will be closed
}