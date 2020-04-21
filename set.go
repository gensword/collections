package collections

import "sync"

type Set struct {
	mut *sync.Mutex
	innerMap     map[interface{}]bool
	results []interface{}
}

func NewSet(elems ...interface{}) *Set {
	set := new(Set)
	set.mut = new(sync.Mutex)
	set.innerMap = make(map[interface{}]bool)
	set.results = make([]interface{}, 0)
	for _, elem := range elems {
		if _, ok := set.innerMap[elem]; !ok {
			set.results = append(set.results, elem)
			set.innerMap[elem] = true
		}
	}
	return set
}

// Add add a elem into set
func (s *Set) Add(elem interface{}){
	s.mut.Lock()
	defer s.mut.Unlock()
	s.results = append(s.results, elem)
	s.innerMap[elem] = true
}

// Exists return true if elem exist else false
func (s *Set) Exists(elem interface{}) bool {
	s.mut.Lock()
	defer s.mut.Unlock()
	_, ok := s.innerMap[elem]
	return ok
}

// Del delete the elem from set
func (s *Set) Del(elem interface{}) bool {
	s.mut.Lock()
	defer s.mut.Unlock()
	if _, ok := s.innerMap[elem]; !ok {
		return false
	} else {
		index, _ := sliceIndex(len(s.results), func(i int) bool { return s.results[i] == elem })
		s.results = append(s.results[:index], s.results[index+1:]...)
		delete(s.innerMap, elem)
		return true
	}
}

//Intersect return the intersection between s and other
func (s *Set) Intersect(other *Set) *Set {
	s.mut.Lock()
	defer s.mut.Unlock()
	ans := NewSet()
	for _, v := range s.results {
		if other.Exists(v) {
			ans.results = append(ans.results, v)
			ans.innerMap[v] = true
		}
	}
	return ans
}

// Union return the union of s and other
func (s *Set) Union(other *Set) *Set {
	s.mut.Lock()
	defer s.mut.Unlock()
	ans := NewSet()
	for _, k := range s.results {
		ans.results = append(ans.results, k)
		ans.innerMap[k] = true
	}
	for _, v := range other.results {
		if !ans.Exists(v) {
			ans.results = append(ans.results, v)
			ans.innerMap[v] = true
		}
	}
	return ans
}

// diff = union - intersect
func (s *Set) Diff(other *Set) *Set {
	s.mut.Lock()
	defer s.mut.Unlock()
	ans := NewSet()
	for _, v := range s.results {
		if _, ok := other.innerMap[v]; !ok {
			ans.results = append(ans.results, v)
			ans.innerMap[v] = true
		}
	}
	for _, ov := range other.results {
		if _, ok := s.innerMap[ov]; !ok {
			ans.results = append(ans.results, ov)
			ans.innerMap[ov] = true
			}
	}
	return ans
}

//Elements return all elems
func (s *Set) Elements() []interface{} {
	s.mut.Lock()
	defer s.mut.Unlock()
	return s.results
}

//Len return length of s
func (s *Set) Len() int {
	s.mut.Lock()
	defer s.mut.Unlock()
	return len(s.results)
}

func sliceIndex(limit int, predicate func(i int) bool) (int, bool) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i, true
		}
	}
	return -1, false
}
