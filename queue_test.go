package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestQueuePushAndPop(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	v, ok := q.Pop()
	assert.Equal(t, 1, v)
	assert.Equal(t, true, ok)
	v, ok = q.Pop()
	assert.Equal(t, nil, v)
	assert.Equal(t, false, ok)
}

func TestQueueSize(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, 0, q.Size())
	q.Push(1)
	assert.Equal(t, 1,  q.Size())
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, true, q.IsEmpty())
	q.Push(1)
	assert.Equal(t, false, q.IsEmpty())
}