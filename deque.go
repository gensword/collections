package collections

import (
	"container/list"
	"sync"
)

type Deque struct {
	data *list.List
	mut  *sync.Mutex
}

func NewDeque() (dq *Deque) {
	dq = new(Deque)
	dq.data = list.New()
	dq.mut = new(sync.Mutex)
	return dq
}

//Append push back a item
func (dq *Deque) Append(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushBack(v)
}

// push front a item
func (dq *Deque) AppendLeft(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushFront(v)
}

//Pop pop a item from back, if dequeue
func (dq *Deque) Pop() (v interface{}, canPop bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		return elem.Value, true
	} else {
		return nil, false
	}
}

//PopLeft pop a item from front
func (dq *Deque) PopLeft() (v interface{}, canGet bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		return elem.Value, true
	} else {
		return nil, false
	}
}

//Clear clear a dequeue
func (dq *Deque) Clear() {
	dq.data = list.New()
}

// Remove remove the first elem match v
func (dq *Deque) Remove(v interface{}) (removed bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() == 0 {
		return false
	}
	find := false
	current := dq.data.Front()
	var needDelete *list.Element
	for i := 0; i < dq.data.Len(); i++ {
		if current.Value == v {
			find = true
			needDelete = current
			break
		} else {
			current = current.Next()
		}
	}
	if find {
		dq.data.Remove(needDelete)
		return true
	} else {
		return false
	}
}

// Index return the first index between start(include) and end(not include) which match v, find will be false if can not find
func (dq *Deque) Index(v interface{}, start int, end int) (position int, find bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if start >= end || start < 0 || start >= dq.data.Len() {
		return 0, false
	} else {
		for current, i := dq.data.Front(), 0; current != nil; current, i = current.Next(), i + 1 {
			if current.Value == v && start <= i && end > i{
				return i, true
			}
		}
		return 0, false
	}
}

//ExtendLeft push front a container/list
func (dq *Deque) ExtendLeft(other *list.List) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushFrontList(other)
}

//Extend push back a container/list
func (dq *Deque) Extend(other *list.List) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushBackList(other)
}

// Rotate right shift step, if step < 0 then left shift step
func (dq *Deque) Rotate(step int) {
	if dq.data.Len() == 0 {
		return
	}
	if step > 0 {
		for ; step > 0; step-- {
			v, _ := dq.Pop()
			dq.AppendLeft(v)
		}
	} else {
		for ; step < 0; step++ {
			v, _ := dq.PopLeft()
			dq.Append(v)
		}
	}
}

// return length of the deque
func (dq *Deque) Size() int {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	return dq.data.Len()
}
