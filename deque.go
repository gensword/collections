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

func (dq *Deque) Append(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushBack(v)
}

func (dq *Deque) AppendLeft(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushFront(v)
}

func (dq *Deque) Pop() (v interface{}, canGet bool) {
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

func (dq *Deque) Clear() {
	for ; dq.data.Len() > 0; dq.Pop(){
	}
}

func (dq *Deque) Remove(v interface{}) (removeNum int, canRemove bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() == 0 {
		return 0, false
	}
	find := false
	current := dq.data.Front()
	var needDelete *list.Element
	for i := 0; i < dq.data.Len(); i++ {
		if current.Value == v {
			if !find {
				find = true
			}
			needDelete = current
			break
		} else {
			current = current.Next()
		}
	}
	if find {
		dq.data.Remove(needDelete)
		return 1, true
	} else {
		return 0, true
	}
}

func (dq *Deque) Index(v interface{}, start int, end int) (position int, find bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if start >= end || start < 0 || start >= dq.data.Len() {
		return -1, false
	} else {
		for current, i := dq.data.Front(), 0; current != nil; current, i = current.Next(), i + 1 {
			if current.Value == v {
				return i, true
			}
		}
		return 0, false
	}
}

func (dq *Deque) ExtendLeft(other *list.List) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushFrontList(other)
}

func (dq *Deque) Extend(other *list.List) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushBackList(other)
}

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

func (dq *Deque) Size() int {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	return dq.data.Len()
}
