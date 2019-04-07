package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	om := newOederedMap()
	om.Set("name", "gensword", true)
	om.Set("age", 21, true)
	v, ok := om.Get("age")
	assert.Equal(t, 21, v)
	assert.Equal(t, true, ok)
	v, ok = om.Get("location")
	assert.Equal(t, false, ok)
	om.Set("name", "hello", false)
	items := []map[interface{}]interface{}{map[interface{}]interface{}{"name": "hello"}, map[interface{}]interface{}{"age": 21}}
	i := 0
	for item := range om.Iter() {
		assert.Equal(t, items[i], item)
		i++
	}
	ok = om.Del("name")
	assert.Equal(t, true, ok)
	ok = om.Del("location")
	assert.Equal(t, false, ok)
	om.Set("location", "shanghai", true)
	om.Set("age", 22, true)
	assert.Equal(t, 22, om.List.tail.pre.value)
}
