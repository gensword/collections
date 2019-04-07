# collections
golang collections like python include ordered map and chain map...
***
implemented *chainmap*, *counter*, *deque(thread safe)*, *queue(thread safe)*, *orderedmap* and *set* datatypes so far.
##Basic usages
- **chain map**
```go
cm := NewChainMap() // empty chain map
child := map[string]interface{}{
		"country": "China",
		"age":     21,
	}
cm.NewChild(child) // add a new map element
other := map[string]interface{}{
            		"name":   "gensword",
            		"gender": "man",
            		"age":    22,
            	}
cm = NewCHainMap(child, other) // get a chain map for child and other like map[string]interface{}{"country": "China", "age": 21, "gender": "man"}
keys := cm.Keys() // return keys for chain map
values := cm.Values() // return values for chain map
parentsChainMap := cm.Parents() // return a new chain map except the first map for cm chain map
parentsChainMap.Maps() // return a map slice
cm.Map["name"] // get gensword
```
- **counter**
```go
counter := NewCounter("a", "b", "c", "d", "b") // init a counter with some elements
counter.Add(1) // add one
counter.Add(1) // now twice 1
top2 := counter.MostCommon(2) // get a slice like PairList{Pair{1, 2}, Pair{"b", 2}} that indicate the top 2 elements and their counts
counter.Elements() //return all elements
counter.Del(1)
len := counter.Len()
```
- **deque**
```go
dq := NewDeque()
dq.AppendLeft(1) // push left
dq.Append("gensword") // push right
dq.Pop() // pop right
dq.PopLeft() // pop left
dq.Extend() // extend a *list.List right
dq.ExtendLeft() // extend a *list.List left
dq.Rotate(1) // right step 1
dq.Rotate(-1) // left step 1
dq.Index(1, 0, dq.Size()) // return the first position index of element and true which value is 1 between 0 and len(dq), if not found, return zero and false
dq.Remove(1) // delete the first element satisfied with 1 and return remove nums(1) and true, if not found, return zero and false
dq.Clear() // clear all elements
```
- **queue**
```go
q := Newqueue()
q.Push(1)
q.push("hello")
q.Pop()
q.Size() // return 1
q.Pop()
q.Isempty // return true
```
- **orderedMap**
```go
om := NewOrderedMap()
om.Set("name", "gensword", true) // set name, the last bool argument is a flag if the name should replace to the last position and update the key when key name already in om, if false, just update the name value
om.Set("age", 21)
v, ok := om.Get("age") // get the value of age, if not found, reutn nil and false
for item := range om.Iter() {
		
	} // the first item will be map[interface{}]interface{}{"name": "gensword"}, and the second will be map[interface{}]interface{}{"age": 21}
om.Del("name") // delete the key name, if not exists, return false
```
- **set**
```go
s := NewSet(1, "hello") // init a set with two elements
s.Add('I') // add one
s.Exists(1) // true
s.Exists(2) // false
other := NewSet("you", "hello")
intersect := s.Intersect(other) // return a new set include "hello"
union := s.Union(other) // return a new set include 1, "hello", "you", 'I'
diff := s.Diff(other) // return a new set include 1 , 'I', and "you"
s.Elemetns // return []interface{}{1, "hello", 'I'}
```
I am new to go. I want to learn go by writing go code. The datatypes above I think are very useful for daily development.
some code inspired by another golang package [collections](https://github.com/chenjiandongx/collections).