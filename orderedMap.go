package collections

type OrderedMap struct {
	Dict map[interface{}]*LinkedNode
	List *LinkedList
}

func NewOederedMap () *OrderedMap {
	om := new(OrderedMap)
	om.Dict = make(map[interface{}]*LinkedNode)
	om.List = NewLinkdList()
	return om
}

func (om *OrderedMap) Set(key interface{}, value interface{}, endIfExists bool) {
	if k, ok := om.Dict[key]; ok {
		if endIfExists {
			om.List.Remove(k)
			om.Dict[key] = om.List.Append(key, value)
		} else {
			om.Dict[key].value = value
		}
	} else {
		om.Dict[key] = om.List.Append(key, value)
	}
}

func (om *OrderedMap) Get(key interface{}) (interface{}, bool){
	if v, ok := om.Dict[key]; ok {
		return v.value, true
	} else {
		return nil, false
	}
}

func (om *OrderedMap) Del(key interface{}) bool {
	if n, ok := om.Dict[key]; ok {
		if ok := om.List.Remove(n); !ok {
			return false
		}
		delete(om.Dict, key)
		return true
	}
	return false
}

func (om *OrderedMap) Iter() chan map[interface{}]interface{} {
	ch := make(chan map[interface{}]interface{})
	go func() {
		for n := range om.List.Iter() {
			ch <- map[interface{}]interface{}{n.key: n.value}
		}
		close(ch)
	}()
	return ch
}
