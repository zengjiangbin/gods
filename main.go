package main

import (
	"fmt"
	"gods/arraystack"
	"gods/linkedstack"
)

func main() {
	// demo1
	fmt.Println("array stack")
	s := arraystack.New()
	s.Push(1).Push(2).Push(3).Push(4)
	for v := s.Pop(); v != nil; v = s.Pop() {
		fmt.Println(v)
	}

	// demo2
	fmt.Println("linked stack")
	ls := linkedstack.New()
	ls.Push(1).Push(2).Push(3).Push(-1)
	for v := ls.Pop(); v != nil; v = ls.Pop() {
		fmt.Println(v)
	}
}
