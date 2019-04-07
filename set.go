package collections

type Set struct {
	Map     map[interface{}]bool
	Results []interface{}
}

func NewSet(elems ...interface{}) *Set {
	set := new(Set)
	set.Map = make(map[interface{}]bool)
	set.Results = make([]interface{}, 0)
	for _, elem := range elems {
		if _, ok := set.Map[elem]; !ok {
			set.Results = append(set.Results, elem)
			set.Map[elem] = true
		}
	}
	return set
}

func (s *Set) Add(elem interface{}) bool {
	if _, ok := s.Map[elem]; !ok {
		s.Results = append(s.Results, elem)
		s.Map[elem] = true
		return true
	} else {
		return false
	}
}

func (s *Set) Exists(elem interface{}) bool {
	_, ok := s.Map[elem]
	return ok
}

func (s *Set) Del(elem interface{}) bool {
	if _, ok := s.Map[elem]; !ok {
		return false
	} else {
		index, _ := sliceIndex(len(s.Results), func(i int) bool { return s.Results[i] == elem })
		s.Results = append(s.Results[:index], s.Results[index+1:]...)
		delete(s.Map, elem)
		return true
	}
}

func (s *Set) Intersect(other *Set) *Set {
	ans := NewSet()
	for _, v := range s.Results {
		if other.Exists(v) {
			ans.Results = append(ans.Results, v)
			ans.Map[v] = true
		}
	}
	return ans
}

func (s *Set) Union(other *Set) *Set {
	ans := NewSet()
	for _, k := range s.Results {
		ans.Results = append(ans.Results, k)
		ans.Map[k] = true
	}
	for _, v := range other.Results {
		if !ans.Exists(v) {
			ans.Results = append(ans.Results, v)
			ans.Map[v] = true
		}
	}
	return ans
}

func (s *Set) Diff(other *Set) *Set {
	ans := NewSet()
	for _, v := range s.Results {
		if !other.Exists(v) {
			ans.Results = append(ans.Results, v)
			ans.Map[v] = true
		}
	}
	for _, ov := range other.Results {
		if !s.Exists(ov) {
			ans.Results = append(ans.Results, ov)
			ans.Map[ov] = true
			}
	}
	return ans
}

func (s *Set) Elements() []interface{} {
	return s.Results
}

func (s *Set) Len() int {
	return len(s.Results)
}

func sliceIndex(limit int, predicate func(i int) bool) (int, bool) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i, true
		}
	}
	return -1, false
}
