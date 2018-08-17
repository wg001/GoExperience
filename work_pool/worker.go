package work_pool

import (
	"fmt"
	"time"
)

type Worker struct {
	id        int64
	job       chan interface{}
	cancel    bool
	isworking bool
}

func NewWorker(id int64) *Worker {
	return &Worker{id: id, cancel: false}
}

func (w *Worker) Dojob() {
	select {
	case <-w.job:
		fmt.Println("get work to do")
	case <-time.After(time.Duration(10*time.Second)):
		fmt.Println("work waited too long")
	}
}
