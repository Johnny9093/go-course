package rate_limiter

import (
	"fmt"
	"net/http"
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

func TestBucketRateLimitedHttp(t *testing.T) {

	// Backend
	backendRps := 1
	burstSize := 4
	tokens := time.Tick(time.Duration(1000 / backendRps) * time.Millisecond)
	bucket := make(chan int, burstSize)

	// Fill bucket
	for i := 0; i < burstSize; i++ {
		bucket <- 1
	}

	go func(tokens <-chan time.Time, bucket chan int) {
		for range tokens {
			bucket <- 1
		}
	}(tokens, bucket)

	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <- bucket:
			w.WriteHeader(200)
			w.Write([]byte("OK"))
		default:
			w.WriteHeader(429)
			w.Write([]byte("Rate Limit Exceeded"))
		}
	})

	http.ListenAndServe(":8080", nil)

	// http.ListenAndServe(":8080", rateLimiterHandler) - all requests go through the handler, regardless of url
}