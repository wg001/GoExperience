package work_pool

import (
	"sync"
	"log"
	"time"
)

type WorkerPool struct {
	isInitialized bool
	worker        []*Worker
	cap           int
	locker        sync.RWMutex
	closed        bool
	maxPoolCount  int
}

func NewPool(maxPoolCount int) *WorkerPool {
	wPool := &WorkerPool{
		isInitialized: true,
		closed:        false,
		cap:           0,
		maxPoolCount:  maxPoolCount,
	}
	return wPool
}

func (p *WorkerPool) GetWorker() *Worker {
	if !p.isInitialized {
		log.Fatal("workpool don't initialed")
	}
	if p.maxPoolCount <= p.cap {
		log.Fatalf("workpool 已经达到最大值了")
	}
	p.locker.Lock()
	defer p.locker.Unlock()
	p.cap += 1
	wworker := NewWorker(time.Now().Unix())
	wworker.Dojob()
	p.worker = append(p.worker, wworker)
	return wworker
}

func (p *WorkerPool) close() {
	p.closed = true
}
