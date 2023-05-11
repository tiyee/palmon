package vo

import "github.com/tiyee/palmon/consts"

type JobMessage struct {
	TaskId  int64            `json:"task_id"`
	State   consts.TaskState `json:"state"`
	Prior   int8             `json:"prior"`
	Payload any              `json:"payload"`
}
type JobResult struct {
	TaskId int64 `json:"task_id"`
	State  int8  `json:"state"`
	Result any   `json:"result"`
}
