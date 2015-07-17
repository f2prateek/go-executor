package executor

type Executor interface {
	enqueue(func())
}
