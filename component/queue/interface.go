package queue

import "github.com/tiyee/palmon/external/storage"

type IPriorQueue interface {
	Push(t storage.Task)
	Pop() storage.Task
	Empty() bool
}
