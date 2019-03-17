package go_routines

import "time"

func main() {
	// use channels to communicate between different go routines
	cn := make(chan int, 10)
	// cn <- 2 // push value to channel

	// run for as go routine
	go func(inputChannel chan int) {
		println("starting goroutine")
		for {
			println("pulling value")
			value := <- inputChannel // since there are no values in the channel, the func is blocked here
			// <- chan (pulling from chanel) is blocking.
			println(value)
		}
	}(cn)
	// exits immediately, we wont see the result

	time.Sleep(3 * time.Second)
	cn <- 13

	time.Sleep(1 * time.Second)

	// main routine ended, all goroutines will be closed
}