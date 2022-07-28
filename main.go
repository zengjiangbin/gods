package main

import (
	"fmt"
	"gods/arraystack"
)

func main() {
	// demo1
	s := arraystack.New()
	s.Push(1).Push(2).Push(3).Push(4)
	for {
		if v := s.Pop(); v != nil {
			fmt.Println(v)
		} else {
			break
		}
	}
}
