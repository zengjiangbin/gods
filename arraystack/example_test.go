package arraystack

import "fmt"

func Example() {
	s := New()
	s.Push(1).Push(2).Push(3).Push(4)
	for v := s.Pop(); v != nil; v = s.Pop() {
		fmt.Println(v)
	}
}
