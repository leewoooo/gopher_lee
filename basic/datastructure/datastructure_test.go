package datastructure

import (
	"container/list"
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	//given
	linkedList := list.New()

	linkedList.PushFront("foo1")
	linkedList.PushBack("foo2")
	linkedList.PushBack("foo3")
	linkedList.PushBack("foo4")

	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
