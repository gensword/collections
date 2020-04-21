package collections

import (
	"sort"
	"sync"
)

type Pair struct {
	Key   interface{} // key add to counter
	Value int         // store counts of the key
}

type PairList []Pair

type Counter struct {
	Count map[interface{}]int
	mut  *sync.Mutex
}

func NewCounter(elems ...interface{}) (c *Counter) {
	initMap := make(map[interface{}]int)
	for _, elem := range elems {
		if v, ok := initMap[elem]; !ok {
			initMap[elem] = 1
		} else {
			initMap[elem] = v + 1
		}
	}
	c = &Counter{
		Count: initMap,
		mut: new(sync.Mutex),
	}
	return
}

// add a elem(key) into the counter and incr the counts of the key
func (c *Counter) Add(elem interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.Count[elem] += 1
}

// return elements of the counter
func (c *Counter) Elements() []interface{} {
	c.mut.Lock()
	defer c.mut.Unlock()
	elements := make([]interface{}, 0)
	for k, v := range c.Count {
		for i := 0; i < v; i++ {
			elements = append(elements, k)
		}
	}
	return elements
}

//return top frequently keys and their counts
func (c *Counter) MostCommon(top int) PairList {
	c.mut.Lock()
	defer c.mut.Unlock()
	if top > len(c.Count) || top < 0 {
		top = len(c.Count)
	}
	p := make(PairList, 0)
	for k, v := range c.Count {
		p = append(p, Pair{k, v})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].Value > p[j].Value
	})
	return p[:top]
}

// delete a key from a counter, if key exist return true else return false
func (c *Counter) Del(key interface{}) bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	if _, ok := c.Count[key]; ok {
		delete(c.Count, key)
		return true
	} else {
		return false
	}
}

// return length of a counter
func (c *Counter) Len() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return len(c.Count)
}
