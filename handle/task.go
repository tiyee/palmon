package handle

import (
	"fmt"
	"github.com/tiyee/palmon/component/db"
	"github.com/tiyee/palmon/component/iden"
	"github.com/tiyee/palmon/component/queue"
	"github.com/tiyee/palmon/consts"
	"github.com/tiyee/palmon/engine"
	"github.com/tiyee/palmon/external/storage"
	"github.com/tiyee/palmon/schema"
	"github.com/tiyee/palmon/vo"
	"strconv"
	"strings"
	"time"
)

func Post(c *engine.Context) {
	var job schema.Job
	if err := c.JSONArgs(&job); err != nil {
		c.AjaxError(1, "json error", err.Error())
		return
	}
	task := storage.Task{
		Id:        iden.Gen(),
		Processor: consts.Processor(job.Processor),
		Payload:   job.Payload,
		Ts:        time.Now().Unix(),
		Prior:     consts.Prior(job.Prior),
		Score:     iden.Score(consts.Prior(job.Prior), 10),
		State:     consts.PENDING,
	}
	db.DB.Save(task)
	queue.PriorQueue.Push(task)
	c.AjaxSuccess("ok", vo.Receipt{TaskId: task.Id})

}
func Get(c *engine.Context) {
	vars := c.Request().URL.Query()
	fmt.Println(c.Request())
	arr, ok := vars["task_id"]
	if !ok {
		c.AjaxError(2, "task_id is missing", 1)
		return
	}
	n, err := strconv.ParseInt(strings.Join(arr, ""), 10, 64)
	if err != nil {
		c.AjaxError(3, "task_id is invalidate", 1)
		return
	}
	fmt.Println(n)
	task, idx := db.DB.Peek(n)
	fmt.Println(idx)
	if idx == -1 {
		c.AjaxError(4, "task not found", 1)
		return
	}
	result := vo.JobResult{TaskId: task.Id, State: int8(task.State), Result: task.Result}
	c.AjaxSuccess("ok", result)
}
