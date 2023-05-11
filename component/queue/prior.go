package queue

import (
	"github.com/tiyee/palmon/external/max_heap"
	"github.com/tiyee/palmon/external/storage"
)

var PriorQueue IPriorQueue

func Init() {
	PriorQueue = max_heap.New[storage.Task]()
}
