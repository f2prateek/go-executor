package executor

import "time"

// Returns an executor that rate limits enqueued tasks.
func NewRateLimitedExecutor(rate time.Duration) Executor {
	e := &rateLimitedExecutor{rate, make(chan func())}
	go e.loop()
	return e
}

type rateLimitedExecutor struct {
	rate  time.Duration
	queue chan func()
}

func (e *rateLimitedExecutor) enqueue(cmd func()) {
	e.queue <- cmd
}

func (e *rateLimitedExecutor) loop() {
	throttle := time.Tick(e.rate)

	for {
		<-throttle
		cmd := <-e.queue
		cmd()
	}
}
