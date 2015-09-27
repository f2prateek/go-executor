package executor

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestPoolExecutor(t *testing.T) {
	e := NewPooledExecutor(2)

	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		e.enqueue(func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			log.Println("hello")
		})
	}

	wg.Wait()
}
