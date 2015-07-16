# go-executor

Experimental executor implementation for Go.

Decouples task submission from mechanics of how each task will be run.

Currently provides a single implemention via the PoolExecutor.

In this example, the pooled executor will guarantee that only upto 20 goroutines run at a given time.
```
func main() {
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
```

Easy to see that in tests you could swap the executor for a simpler implementation, such as a syncronous executor.