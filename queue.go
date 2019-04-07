package collections

import (
	"container/list"
	"sync"
)

type Queue struct {
	data *list.List
	mut  *sync.Mutex
}

func NewQueue() (queue *Queue) {
	queue = new(Queue)
	queue.data = list.New()
	queue.mut = new(sync.Mutex)
	return
}

func (queue *Queue) Push(v interface{}) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	queue.data.PushFront(v)
}

func (queue *Queue) IsEmpty() (isEmpty bool) {
	if queue.data.Len() > 0 {
		return false
	} else {
		return true
	}
}

func (queue *Queue) Pop() (v interface{}, canGet bool) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return nil, false
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		return elem.Value, true
	}
}

func (queue *Queue) Size() int {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	return queue.data.Len()
}
