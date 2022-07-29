package arrayqueue

import "fmt"

func Example() {
	aq := New()
	aq.Push(1).Push(2).Push(3).Push(4)
	for v := aq.Pop(); v != nil; v = aq.Pop() {
		fmt.Println(v)
	}
}
