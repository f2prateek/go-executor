package executor

// Returns an executor that runs all tasks sequentially.
func NewSequentialExecutor() Executor {
	return NewPooledExecutor(1)
}

func NewPooledExecutor(poolCount int) Executor {
	p := &poolExecutor{poolCount, make(chan func())}
	go p.loop()
	return p
}

type poolExecutor struct {
	poolCount int
	queue     chan func()
}

func (p *poolExecutor) enqueue(cmd func()) {
	p.queue <- cmd
}

func (p *poolExecutor) loop() {
	var q []func()
	running := 0
	done := make(chan bool)

	for {
		select {
		case cmd := <-p.queue:
			if running == p.poolCount {
				q = append(q, cmd)
				continue
			}

			running++
			go func() {
				cmd()
				done <- true
			}()
		case <-done:
			running--
			if len(q) == 0 {
				continue
			}

			cmd := q[0]
			q = q[1:]
			running++

			go func() {
				cmd()
				done <- false
			}()
		}
	}
}
