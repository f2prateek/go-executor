package executor

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolExecutor(t *testing.T) {
	e := NewPooledExecutor(20)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		e.enqueue(func() {
			defer wg.Done()
			time.Sleep(500 * time.Millisecond)
			fmt.Println("hello")
		})
	}

	wg.Wait()
}
