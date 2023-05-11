package comsumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiyee/palmon/consts"
	"github.com/tiyee/palmon/lib"
	"github.com/tiyee/palmon/schema"
	"github.com/tiyee/palmon/vo"
	"io"
	"log"
	"net/http"
	"time"
)

type Puller struct {
	worker *worker
}

func (p Puller) Run(worker *worker, job Job) {
	url := "http://coordinator:8083/sync"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("pull error: ", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		log.Println("pull error: code=", resp.Status)
		return
	}
	defer resp.Body.Close()
	var msg schema.SyncJob
	if err := lib.JSONArgs(resp.Body, &msg); err != nil {
		log.Println("pull data parse error, error=", err.Error())
	}
	if msg.Error != 0 {
		log.Println("pull error: code=", msg.Error, " message=", msg.Message)
		return
	}
	work := Job{Payload: msg.Data, Processor: consts.PROCESSOR1}
	// Push the work onto the queue.
	worker.dispatcher.jobQueue <- work
}
func (p Puller) Name() string {
	return "puller"
}

type Pusher struct {
}

func (p Pusher) Run(worker *worker, job Job) {
	url := "http://coordinator:8083/sync"
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(job.Payload))
	if err != nil {
		log.Println("push error: ", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		log.Println("push error: code=", resp.Status)
		return
	}
	defer resp.Body.Close()
	if bs, err := io.ReadAll(resp.Body); err == nil {
		log.Println("push success: ", string(bs))
	} else {
		log.Println("push error: ", err.Error())
	}

}
func (p Pusher) Name() string {
	return "pusher"
}

type Processor1 struct {
	worker  *worker
	payload json.RawMessage
}

func (p Processor1) Run(worker *worker, job Job) {
	var message vo.JobMessage
	if err := json.Unmarshal(job.Payload, &message); err != nil {
		log.Printf("proc data parse error, processor=%s error=%s payload=%s\n", p.Name(), err.Error(), job.Payload)
		return
	}

	result := schema.TaskResult{
		TaskId: message.TaskId,
		State:  consts.FULFILL,
		Result: fmt.Sprint(message.Payload, message.Payload),
	}
	bs, err := json.Marshal(result)
	if err != nil {
		log.Printf("proc data encode error, processor=%s error=%s\n", p.Name(), err.Error())
		return
	}
	work := Job{Payload: bs, Processor: consts.PUSHER}
	rand := time.Now().UnixNano() % 9999
	// rand delay
	time.Sleep(time.Millisecond * time.Duration(rand))

	// Push the work onto the queue.
	worker.dispatcher.jobQueue <- work

}
func (p Processor1) Name() string {
	return "proc1"
}

func Claim(worker *worker, job Job) {
	if proc, exist := processors[job.Processor]; exist {
		proc.Run(worker, job)
		fmt.Println("run: ", proc.Name(), " ", string(job.Payload))

	} else {
		fmt.Printf("proc %d is not found \n", job.Processor)
	}
}
