package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s := NewSet(1, "hello")
	s.Add('I')
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, true, s.Exists('I'))
	assert.Equal(t, false, s.Exists("you"))
	s.Del(1)
	assert.Equal(t, false, s.Exists(1))
	assert.Equal(t, 2, s.Len())
	other := NewSet("you", "hello")
	intersect := s.Intersect(other)
	assert.Equal(t, 1, intersect.Len())
	assert.Equal(t, "hello", intersect.Elements()[0])
	diff := s.Diff(other)
	assert.Equal(t, []interface{}{'I', "you"}, diff.Elements())
	union := s.Union(other)
	assert.Equal(t, []interface{}{"hello", 'I', "you"}, union.Elements())
}
