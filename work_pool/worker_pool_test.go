package work_pool

import (
	"testing"
	"time"
)

func TestNewWorker(t *testing.T) {//不行,还要再设计
	MaxPoolCount := 10
	pool := NewPool(MaxPoolCount)
	for i := 0; i < MaxPoolCount; i++ { //应该是从池中获取工人，然后一个工人做完任务后再不断的接收新的任务，做任务
		go pool.GetWorker()
	}
	time.Sleep(15*time.Second)
}
