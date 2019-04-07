package collections

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestChainMap(t *testing.T) {
	cm := NewChainMap()
	child := map[string]interface{}{
		"country": "China",
		"age":     21,
	}
	other := map[string]interface{}{
		"name":   "gensword",
		"gender": "man",
		"age":    22,
	}
	cm.NewChild(child)
	assert.Contains(t, [][]string{[]string{"country", "age"}, []string{"age", "country"}}, cm.Keys())
	assert.Contains(t, [][]interface{}{[]interface{}{"China", 21}, []interface{}{21, "China"}}, cm.Values())
	cm = NewChainMap(child, other)
	assert.Equal(t, 4, len(cm.Map))
	assert.Equal(t, 21, cm.Map["age"])
	assert.True(t, reflect.DeepEqual([]map[string]interface{}{map[string]interface{}{"name": "gensword",
		"gender": "man",
		"age":    22,}}, cm.Parents().Maps))
}
