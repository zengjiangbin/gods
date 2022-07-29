// Package linkedstack  实现了非线程安全的链栈
package linkedstack

import "container/list"

// 栈结构
type LinkedStack struct {
	data *list.List
}

// New 创建栈
func New() *LinkedStack {
	return new(LinkedStack).init()
}

// init 初始化栈
func (l *LinkedStack) init() *LinkedStack {
	l.data = list.New()
	return l
}

// Len 返回栈长度
func (l *LinkedStack) len() int {
	return l.data.Len()
}

// Clear 清空栈
func (l *LinkedStack) Clear() *LinkedStack {
	return l.init()
}

// Push 入栈
func (l *LinkedStack) Push(v interface{}) *LinkedStack {
	l.data.PushFront(v)
	return l
}

// Pop 出栈
func (l *LinkedStack) Pop() interface{} {
	if l.len() < 1 {
		return nil
	}
	return l.data.Remove(l.data.Front())
}
