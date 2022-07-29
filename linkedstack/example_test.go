package linkedstack

import "fmt"

func Example() {
	ls := New()
	ls.Push(1).Push(2).Push(3).Push(-1)
	for v := ls.Pop(); v != nil; v = ls.Pop() {
		fmt.Println(v)
	}
}
