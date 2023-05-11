package handle

import (
	"github.com/tiyee/palmon/component/db"
	"github.com/tiyee/palmon/component/queue"
	"github.com/tiyee/palmon/engine"
	"github.com/tiyee/palmon/schema"
	"github.com/tiyee/palmon/vo"
)

func Pull(c *engine.Context) {
	if queue.PriorQueue.Empty() {
		c.AjaxError(5, "task queue is empty", 0)
		return
	}
	task := queue.PriorQueue.Pop()
	jobMessage := vo.JobMessage{
		TaskId:  task.Id,
		State:   task.State,
		Payload: task.Payload,
		Prior:   int8(task.Prior),
	}
	c.AjaxSuccess("ok", jobMessage)
}
func Push(c *engine.Context) {
	var job schema.TaskResult
	if err := c.JSONArgs(&job); err != nil {
		c.AjaxError(1, "json error", 1)
		return
	}
	task, idx := db.DB.Peek(job.TaskId)
	if idx == -1 {
		c.AjaxError(4, "task not found", 1)
		return
	}
	task.Result = job.Result
	task.State = job.State
	db.DB.Set(idx, task)
	c.AjaxSuccess("ok", idx)
}
