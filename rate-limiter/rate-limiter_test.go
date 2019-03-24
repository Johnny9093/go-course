package rate_limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	Limit()
	time.Sleep(5 * time.Second)
}

func TestHttpLimiterAndMock(t *testing.T) {
	network := make(chan string, 20)

	// Client Mock
	go func(netLine chan string) {
		clientRate := 10
		index := 0
		for {
			network <- fmt.Sprintf("This is a mock request: %d", index)
			time.Sleep(time.Duration(1000/clientRate) * time.Millisecond)
			index++
		}
	}(network)

	// Backend
	backendRps := 2
	ticker := time.Tick(time.Duration(1000 / backendRps) * time.Millisecond)
	go func(netLine chan string) {
		for request := range netLine {
			select {
			case <- ticker:
				fmt.Printf("200 OK: %s\n", request)
			default:
				fmt.Printf("429 Rate Limit Exceeded: %s\n", request)
			}
		}
	}(network)

	time.Sleep(5 * time.Second)
}

func TestBucketTokenRateLimiter(t *testing.T) {
	network := make(chan string, 20)

	// Client Mock
	go func(netLine chan string) {
		clientRate := 20
		index := 0
		for {
			network <- fmt.Sprintf("This is a mock request: %d", index)
			time.Sleep(time.Duration(1000/clientRate) * time.Millisecond)
			index++
		}
	}(network)

	// Backend
	backendRps := 2
	burstSize := 10
	tokens := time.Tick(time.Duration(1000 / backendRps) * time.Millisecond)
	bucket := make(chan int, burstSize)

	// Fill bucket
	for i := 0; i < burstSize; i++ {
		bucket <- 1
	}

	go func(netLine chan string) {
		for request := range netLine {
			select {
			case <- bucket:
				fmt.Printf("200 OK: %s\n", request)
			default:
				fmt.Printf("429 Rate Limit Exceeded: %s\n", request)
			}
		}
	}(network)

	go func(tokens <-chan time.Time, bucket chan int) {
		for {
			<- tokens
			bucket <- 1
		}
	}(tokens, bucket)

	time.Sleep(5 * time.Second)
}