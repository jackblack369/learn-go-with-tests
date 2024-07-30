package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer compares the response times of a and b, returning the fastest one.
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	fmt.Println("aDuration", aDuration)
	bDuration := measureResponseTime(b)
	fmt.Println("bDuration", bDuration)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
