package work_pool

import (
	"fmt"
	"strconv"
)

type Worker struct {
	id        int64
	job       interface{}
	cancel    bool
	isworking bool
}

func NewWorker(id int64) *Worker {
	return &Worker{id: id, cancel: false}
}

func (w *Worker) Dojob() {
	fmt.Println("工人[id=" + strconv.Itoa(int(w.id)) + "]在做工作")
}
