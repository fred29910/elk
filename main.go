package main

import (
	"context"
	"fmt"
	"time"
)

// MyContext 是自定义的 context
type MyContext struct {
	context.Context               // 嵌入标准 context，以便继承其行为
	done            chan struct{} // 用于通知取消的通道
	err             error         // 保存取消时的错误
}

// NewMyContext 创建一个自定义的 context
func NewMyContext(parent context.Context) *MyContext {
	return &MyContext{
		Context: parent,
		done:    make(chan struct{}),
	}
}

// Done 返回一个通道，当操作应该取消时关闭这个通道
func (c *MyContext) Done() <-chan struct{} {
	return c.done
}

// Err 返回 context 被取消的原因
func (c *MyContext) Err() error {
	return c.err
}

// Cancel 手动取消 context
func (c *MyContext) Cancel(err error) {
	c.err = err
	close(c.done)
}

// Deadline 实现 context.Context 的 Deadline 方法
func (c *MyContext) Deadline() (deadline time.Time, ok bool) {
	deadline, ok = c.Context.Deadline()
	return
}

// Value 实现 context.Context 的 Value 方法
func (c *MyContext) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

func main() {
	// 创建一个父 context
	parentCtx := context.Background()

	// 创建一个自定义的 MyContext
	myCtx := NewMyContext(parentCtx)

	// 启动一个 goroutine 来监听自定义 context
	go func() {
		select {
		case <-myCtx.Done():
			fmt.Println("MyContext cancelled:", myCtx.Err())
			return
		}
	}()

	// 模拟工作
	time.Sleep(2 * time.Second)

	// 手动取消自定义 context
	myCtx.Cancel(fmt.Errorf("operation timed out"))

	// 等待 goroutine 处理取消信号
	time.Sleep(1 * time.Second)
}
