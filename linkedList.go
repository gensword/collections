// Copyright tizz98 The Go Authors. All rights reserved.
// Github https://github.com/tizz98/go-playground/tree/master/linkedlist

package collections

type LinkedNode struct {
	pre   *LinkedNode
	next  *LinkedNode
	key interface{}
	value interface{}
}

type LinkedList struct {
	head *LinkedNode
	tail *LinkedNode
}

func NewLinkdList() *LinkedList {
	list := new(LinkedList)
	list.head, list.tail = &LinkedNode{key: nil, value: nil, pre: nil, next: nil}, &LinkedNode{key: nil, value: nil, pre: nil, next: nil}
	list.head.next, list.tail.pre = list.tail, list.head
	return list
}

func (l *LinkedList) Append(key interface{}, value interface{}) *LinkedNode {
	newNode := &LinkedNode{pre: l.tail.pre, next: l.tail, key: key, value: value}
	newNode.pre.next = newNode
	l.tail.pre = newNode
	return newNode
}

func (l *LinkedList) Remove(n *LinkedNode) bool {
	if n == l.head || n == l.tail {
		return false
	}

	n.pre.next = n.next
	n.next.pre = n.pre

	return true
}

func (l *LinkedList) Iter() chan *LinkedNode {
	ch := make(chan *LinkedNode)
	go func() {
		for current := l.head; current.next != l.tail; current = current.next {
			ch <- current.next
		}
		close(ch)
	}()
	return ch
}
