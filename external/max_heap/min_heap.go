package max_heap

import (
	"sort"
	"sync"
)

type Comparable interface {
	CmpValue() int64
}
type Heap[T Comparable] []T

func (h *Heap[T]) Len() int {
	return len(*h)
}

func (h *Heap[T]) Less(i, j int) bool {
	return (*h)[i].CmpValue() < (*h)[j].CmpValue()
}

func (h *Heap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
	sort.Sort(h)
}
func (h *Heap[T]) Pop() T {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func New[T Comparable]() *PriorQueue[T] {

	return &PriorQueue[T]{queue: &Heap[T]{}}
}

type PriorQueue[T Comparable] struct {
	queue *Heap[T]
	lock  sync.Mutex
}

func (h *PriorQueue[T]) Push(x T) {
	h.lock.Lock()

	defer h.lock.Unlock()
	h.queue.Push(x)
}
func (h *PriorQueue[T]) Pop() T {
	h.lock.Lock()

	defer h.lock.Unlock()
	return h.queue.Pop()
}
func (h *PriorQueue[T]) Empty() bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(*h.queue) == 0
}
