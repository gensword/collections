package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque (t *testing.T) {
	dq := NewDeque()
	waitAppend := []interface{}{1,2,3,4,5}
	for _, v := range waitAppend {
		dq.Append(v)
	}
	index, find := dq.Index(waitAppend[1], 0, len(waitAppend))
	assert.Equal(t, 1, index)
	assert.Equal(t, true, find)
	index, find = dq.Index(waitAppend[1], 2, len(waitAppend))
	assert.Equal(t, 0, index)
	assert.Equal(t, false, find)
	size := dq.Size()
	assert.Equal(t, size, 5)
	for i := 0; i <= size; i++ {
		v, canPop := dq.PopLeft()
		if i == size {
			assert.Equal(t, false, canPop)
			assert.Equal(t, nil, v)
		} else {
			assert.Equal(t, true, canPop)
			assert.Equal(t, waitAppend[i], v)
		}
	}
	assert.Equal(t, dq.Size(), 0)
	for _, v := range waitAppend {
		dq.AppendLeft(v)
	}
	size = dq.Size()
	for i := 0; i <= size; i++ {
		v, canPop := dq.Pop()
		if i == size {
			assert.Equal(t, false, canPop)
			assert.Equal(t, nil, v)
		} else {
			assert.Equal(t, true, canPop)
			assert.Equal(t, waitAppend[i], v)
		}
	}
	assert.Equal(t, false, dq.Remove(waitAppend[0]))
	for _, v := range waitAppend {
		dq.Append(v)
	}
	step := 2
	dq.Rotate(step)
	for i := 0; i < len(waitAppend); i++{
		var pos int
		if i + step >= len(waitAppend) {
			pos = i + step - len(waitAppend)
		} else {
			pos = i + step
		}
		actual, _ := dq.Index(waitAppend[i], 0, len(waitAppend))
		assert.Equal(t, pos, actual)
	}
	dq.Clear()
	assert.Equal(t, 0, dq.Size())
}
