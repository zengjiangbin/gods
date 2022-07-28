// Package arraystack 实现了非线程安全的顺序栈
// 使用姿势看 example_test.go
package arraystack

// 栈结构
type arrayStack struct {
	data []interface{}
	top  int
}

// init 初始化栈
func (a *arrayStack) init() *arrayStack {
	a.data = make([]interface{}, 0)
	a.top = -1
	return a
}

// New 创建栈
func New() *arrayStack {
	return new(arrayStack).init()
}

// Len 返回栈长度
func (a *arrayStack) Len() int {
	return a.top + 1
}

// Push 入栈
func (a *arrayStack) Push(v interface{}) *arrayStack {
	a.data = append(a.data, v)
	a.top++
	return a
}

// Pop 出栈
func (a *arrayStack) Pop() interface{} {
	if a.top == -1 {
		return nil
	}
	v := a.data[a.top]
	a.data = a.data[:a.top]
	a.top--
	return v
}

// Clear 清空栈
func (a *arrayStack) Clear() *arrayStack {
	return a.init()
}
