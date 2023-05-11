package comsumer

import (
	"fmt"
	"github.com/tiyee/palmon/consts"
	"strconv"
	"time"
)

func Add() {
	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {
		for { //循环
			<-ticker.C
			fmt.Println("i = ")
		}
	}()
}

type Dispatcher struct {
	workerPool chan *worker
	jobQueue   chan Job
	quit, exit chan bool
	maxWorkers int
}

func NewDispatcher(maxWorkers, maxJobs int) *Dispatcher {

	return &Dispatcher{
		workerPool: make(chan *worker, maxWorkers),
		jobQueue:   make(chan Job, maxJobs),
		quit:       make(chan bool),
		exit:       make(chan bool),
		maxWorkers: maxWorkers,
	}
}
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := newWorker(d, d.workerPool)

		worker.start()
	}

	go d.dispatch()
	i := 0
	for {
		i++
		time.Sleep(time.Second * 1)
		d.jobQueue <- Job{Processor: consts.PULLER, Payload: []byte(strconv.FormatInt(int64(i), 10))}
		if i == 1000000 {
			d.quit <- true
			time.Sleep(time.Second * 3)
			break
		}

	}
	fmt.Println("stop ", i)

}
func (d *Dispatcher) stop() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := <-d.workerPool
		worker.stop <- true
		<-worker.stop
	}
	fmt.Println("dispatch stop ")

}
func (d *Dispatcher) dispatch() {

	for {
		select {
		case job := <-d.jobQueue:
			worker := <-d.workerPool
			worker.jobChannel <- job
		case <-d.quit:
			d.stop()
			d.quit <- true
			return
		}
	}

}
