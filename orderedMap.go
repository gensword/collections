package collections

import "sync"

type OrderedMap struct {
	mut *sync.Mutex
	dict map[interface{}]*LinkedNode
	list *LinkedList
}

type item struct {
	Key 	interface{}
	Value 	interface{}
}

func NewOrderedMap () *OrderedMap {
	om := new(OrderedMap)
	om.dict = make(map[interface{}]*LinkedNode)
	om.list = NewLinkdList()
	om.mut = new(sync.Mutex)
	return om
}

// Set set k-v and push back the k-v, the order will not change if the key already exist unless putEndIfExists is true
func (om *OrderedMap) Set(key interface{}, value interface{}, putEndIfExists bool) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if k, ok := om.dict[key]; ok {
		if putEndIfExists {
			om.list.Remove(k)
			om.dict[key] = om.list.Append(key, value)
		} else {
			om.dict[key].value = value
		}
	} else {
		om.dict[key] = om.list.Append(key, value)
	}
}

// Get
func (om *OrderedMap) Get(key interface{}) (interface{}, bool){
	om.mut.Lock()
	defer om.mut.Unlock()
	if v, ok := om.dict[key]; ok {
		return v.value, true
	} else {
		return nil, false
	}
}

// Del del a key from ordered map, return false if key not exist
func (om *OrderedMap) Del(key interface{}) bool {
	om.mut.Lock()
	defer om.mut.Unlock()
	if n, ok := om.dict[key]; ok {
		if ok := om.list.Remove(n); !ok {
			return false
		}
		delete(om.dict, key)
		return true
	}
	return false
}

// return k-v in order(set order)
func (om *OrderedMap) Iter() chan *item {
	om.mut.Lock()
	defer om.mut.Unlock()
	ch := make(chan *item)
	go func() {
		for n := range om.list.Iter() {
			ch <- &item{
				Key: 	n.key,
				Value:	n.value,
			}
		}
		close(ch)
	}()
	return ch
}
