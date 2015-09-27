package executor

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestRateLimitedExecutor(t *testing.T) {
	e := NewRateLimitedExecutor(time.Second * 1)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		e.enqueue(func() {
			defer wg.Done()
			log.Println("hello")
		})
	}

	wg.Wait()
}
