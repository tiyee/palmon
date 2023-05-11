package schema

import "github.com/tiyee/palmon/consts"

type TaskResult struct {
	TaskId int64            `json:"task_id"`
	State  consts.TaskState `json:"state"`
	Result string           `json:"payload"`
}

func (r *TaskResult) Valid() error {
	return nil
}
func (r *TaskResult) Hook() {

}
