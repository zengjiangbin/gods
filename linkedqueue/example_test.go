package linkedqueue

import "fmt"

func Example() {
	lq := New()
	lq.Push(1).Push(2).Push(3).Push(-1)
	for v := lq.Pop(); v != nil; v = lq.Pop() {
		fmt.Println(v)
	}
}
