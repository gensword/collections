package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	om := NewOrderedMap()
	om.Set("name", "gensword", true)
	om.Set("age", 21, true)
	om.Set("gender", 1, true)
	v, ok := om.Get("age")
	assert.Equal(t, 21, v)
	assert.Equal(t, true, ok)
	v, ok = om.Get("location")
	assert.Equal(t, false, ok)
	om.Set("age", 22, false)
	om.Set("name", "hello", true)
	i := 0
	orderedKey := []interface{}{"age", "gender", "name"}
	orderedValue := []interface{}{22, 1, "hello"}
	for item := range om.Iter() {
		assert.Equal(t, orderedKey[i], item.Key)
		assert.Equal(t, orderedValue[i], item.Value)
		i++
	}
	ok = om.Del("name")
	assert.Equal(t, true, ok)
	ok = om.Del("location")
	assert.Equal(t, false, ok)
}
