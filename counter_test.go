package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := NewCounter("a", "b", "c", "d", "b")
	assert.Equal(t, 4, counter.Len())
	counter.Add("e")
	assert.Equal(t, 5, counter.Len())
	elms := counter.Elements()
	assert.Equal(t, 6, len(elms))
	counter.Add(1)
	counter.Add(1)
	top2 := counter.MostCommon(2)
	assert.Contains(t, []PairList{PairList{Pair{1, 2}, Pair{"b", 2}}, PairList{Pair{"b", 2}, Pair{1, 2}}}, top2)
	counter.Del(1)
	assert.Equal(t, 5, counter.Len())
}