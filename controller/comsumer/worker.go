package comsumer

import (
	"encoding/json"
	"github.com/tiyee/palmon/consts"
)

type Job struct {
	Processor consts.Processor
	Payload   json.RawMessage
}
type worker struct {
	workerPool chan *worker
	jobChannel chan Job
	stop       chan bool
	dispatcher *Dispatcher
}

func newWorker(dispatcher *Dispatcher, pool chan *worker) *worker {
	return &worker{
		workerPool: pool,
		dispatcher: dispatcher,
		jobChannel: make(chan Job),
		stop:       make(chan bool),
	}
}
func (w *worker) start() {
	go func() {
		var job Job
		for {
			w.workerPool <- w
			select {
			case job = <-w.jobChannel:
				Claim(w, job)
			case <-w.stop:
				w.stop <- true
				return
			}
		}
	}()

}
