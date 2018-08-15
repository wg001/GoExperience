package work_pool

import "testing"

func TestNewWorker(t *testing.T) {
	MaxPoolCount := 10
	pool := NewPool(MaxPoolCount)
	for i := 0; i < MaxPoolCount; i++ { //应该是从池中获取工人，然后一个工人做完任务后再不断的接收新的任务，做任务
		newWorker := NewWorker(int64(i + 1))
		pool.AddWorker(newWorker)
	}
}
