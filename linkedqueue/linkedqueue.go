// Package linkedqueue 实现了非线程安全的链表队列
// 使用姿势看 example_test.go
package linkedqueue

import (
	"container/list"
)

// 队列结构
type LinkedQueue struct {
	data *list.List
}

// New 创建队列
func New() *LinkedQueue {
	return new(LinkedQueue).init()
}

// init 初始化
func (l *LinkedQueue) init() *LinkedQueue {
	l.data = list.New()
	return l
}

// Len 返回长度
func (l *LinkedQueue) Len() int {
	return l.data.Len()
}

// Clear 清空
func (l *LinkedQueue) Clear() *LinkedQueue {
	l.data.Init()
	return l
}

// Push 入队
func (l *LinkedQueue) Push(v interface{}) *LinkedQueue {
	l.data.PushBack(v)
	return l
}

// Pop 出队
func (l *LinkedQueue) Pop() interface{} {
	if l.Len() == 0 {
		return nil
	}
	return l.data.Remove(l.data.Front())
}
