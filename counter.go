package collections

import (
	"sort"
)

type Pair struct {
	Key   interface{}
	Value int
}

type PairList []Pair

type Counter struct {
	Count map[interface{}]int
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
	}
	return
}

func (c *Counter) Add(elem interface{}) {
	if v, ok := c.Count[elem]; !ok {
		c.Count[elem] = 1
	} else {
		c.Count[elem] = v + 1
	}
}

func (c *Counter) Elements() []interface{} {
	elements := make([]interface{}, 0)
	for k, v := range c.Count {
		for i := 0; i < v; i++ {
			elements = append(elements, k)
		}
	}
	return elements
}

func (c *Counter) MostCommon(top int) PairList {
	if top > c.Len() || top < 0 {
		top = c.Len()
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

func (c *Counter) Del(key interface{}) bool {
	if _, ok := c.Count[key]; ok {
		delete(c.Count, key)
		return true
	} else {
		return false
	}
}

func (c *Counter) Len() int {
	return len(c.Count)
}
