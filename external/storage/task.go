package storage

import (
	"errors"
	"fmt"
	"github.com/tiyee/palmon/consts"
	"sync"
)

type Task struct {
	Id        int64
	Processor consts.Processor
	Payload   any
	Ts        int64
	Score     int64
	Prior     consts.Prior
	State     consts.TaskState
	Result    any
}

func (t Task) CmpValue() int64 {
	return t.Score
}

type Tasks struct {
	tasks []Task
	lock  sync.Mutex
}

func (t *Tasks) search(id int64) (Task, int) {
	t.lock.Lock()

	defer t.lock.Unlock()
	left := 0
	right := len(t.tasks) - 1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Println("mid=", mid, t.tasks)
		if t.tasks[mid].Id == id {
			return t.tasks[mid], mid
		}
		if t.tasks[mid].Id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return Task{}, -1
}
func (t *Tasks) Peek(id int64) (Task, int) {
	return t.search(id)
}
func (t *Tasks) Set(idx int, task Task) {
	t.lock.Lock()

	defer t.lock.Unlock()
	t.tasks[idx] = task
}
func (t *Tasks) Save(task Task) {
	t.lock.Lock()

	defer t.lock.Unlock()
	t.tasks = append(t.tasks, task)
}
func (t *Tasks) Update(task Task) error {
	_, idx := t.search(task.Id)
	if idx == -1 {
		return errors.New("not found")
	}
	t.Set(idx, task)
	return nil
}
