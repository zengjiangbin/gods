// Package arrayqueue 实现了非线程安全的顺序队列
package arrayqueue

// 队列结构
type ArrayQueue struct {
	data []interface{}
}

// New 创建队列
func New() *ArrayQueue {
	return new(ArrayQueue).Init()
}

// Init 初始化队列
func (a *ArrayQueue) Init() *ArrayQueue {
	a.data = make([]interface{}, 0)
	return a
}

// Clear 清理队列
func (a *ArrayQueue) Clear() *ArrayQueue {
	return a.Init()
}

// Len 返回队列长度
func (a *ArrayQueue) Len() int {
	return len(a.data)
}

// Push 入队
func (a *ArrayQueue) Push(v interface{}) *ArrayQueue {
	a.data = append(a.data, v)
	return a
}

// Pop 出队
func (a *ArrayQueue) Pop() interface{} {
	if a.Len() == 0 {
		return nil
	}
	v := a.data[0]
	a.data = a.data[1:]
	return v
}
