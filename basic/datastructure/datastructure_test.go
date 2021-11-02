package datastructure

import (
	"container/list"
	"container/ring"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//given
	l := list.New()

	// 3 -> 4 -> 1 -> 2
	e4 := l.PushBack(4)
	e1 := l.PushBack(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e.Value)
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		t.Log(e.Value)
	}
}

type Queue struct {
	l *list.List
}

func NewQueue(l *list.List) *Queue {
	return &Queue{l: l}
}

func (q *Queue) Push(val interface{}) {
	q.l.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	element := q.l.Front()

	if element != nil {
		return q.l.Remove(element)
	}

	return nil
}

func TestQueue(t *testing.T) {
	queue := NewQueue(list.New())

	// 1 -> 2 -> 3 ->4
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)

	// 1 -> 2 -> 3 -> 4
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
}

type Stack struct {
	l *list.List
}

func NewStack(l *list.List) *Stack {
	return &Stack{l: l}
}

func (s *Stack) Push(val interface{}) {
	s.l.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	element := s.l.Back()

	if element != nil {
		return s.l.Remove(element)
	}

	return nil
}

type Stack2 struct {
	l []interface{}
}

func NewStack2(l []interface{}) *Stack2 {
	return &Stack2{l: l}
}

func (s *Stack2) Push(val interface{}) {
	_ = append(s.l, val)
}

func (s *Stack2) Pop() interface{} {
	last := s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return last
}

func TestStack(t *testing.T) {
	stack := NewStack(list.New())

	// 1 -> 2 -> 3 -> 4
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	// 4 -> 3 -> 2 -> 1
	val := stack.Pop()
	for val != nil {
		t.Log(val)
		val = stack.Pop()
	}
}

func TestRing(t *testing.T) {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		t.Logf("%c", r.Value)
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		t.Logf("%c", r.Value)
		r = r.Prev()
	}
}
