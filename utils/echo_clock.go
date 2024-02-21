package utils

import (
	"sync"
	"time"

	"github.com/golang/glog"
)

var (
	counter int
	mu      sync.Mutex
)

// StartLogging starts a goroutine that logs the count of data pushes every specified interval.
func EchoClock(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			glog.Infof("Number of data pushed in the last %v seconds: %d", interval.Seconds(), counter)
			resetCounter()
		}
	}()
}

// IncrementCounter safely increments the global counter.
func IncrementCounter() {
	mu.Lock()
	counter++
	mu.Unlock()
}

// ResetCounter safely resets the global counter to zero.
func resetCounter() {
	mu.Lock()
	counter = 0
	mu.Unlock()
}
