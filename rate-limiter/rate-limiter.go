package rate_limiter

import (
	"fmt"
	"time"
)

func Limit() {
	client := time.Tick(700 * time.Microsecond)
	server := make(chan time.Time, 10)

	go func(cl <-chan time.Time, sv chan time.Time){
		for {
			time := <-cl
			sv <- time
		}
	}(client, server)

	go func(sv chan time.Time) {
		req := <- sv
		fmt.Printf("%+v", req)
	}(server)
}