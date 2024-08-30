package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// SemaphoreContext 是一个自定义的 Context，用于处理信号量
type SemaphoreContext struct {
	context.Context
	cancelFunc context.CancelFunc
	err        error
	mu         sync.Mutex
}

// NewSemaphoreContext 创建一个带有信号量取消功能的 context
func NewSemaphoreContext(parent context.Context, tm time.Duration) (*SemaphoreContext, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(tm))
	return &SemaphoreContext{
		Context:    ctx,
		cancelFunc: cancel,
	}, cancel
}

// CancelWithError 取消信号量并传递错误
func (s *SemaphoreContext) CancelWithError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.err = err
	s.cancelFunc()
}

// Done 实现 context.Context 的 Done 方法，返回信号量的通道
func (s *SemaphoreContext) Done() <-chan struct{} {
	return s.Context.Done()
}

// Value 实现 context.Context 的 Value 方法，返回自定义值
func (s *SemaphoreContext) Value(key any) any {
	return s.Context.Value(key)
}

// Cancel 实现 context.Context 的 Cancel 方法，取消信号量
func (s *SemaphoreContext) Cancel() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cancelFunc()
}

// Err 实现 context.Context 的 Err 方法，返回自定义错误
func (s *SemaphoreContext) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return s.err
	}
	return s.Context.Err()
}

// Deadline 实现 context.Context 的 Deadline 方法，返回超时时间
func (s *SemaphoreContext) Deadline() (deadline time.Time, ok bool) {
	return s.Context.Deadline()
}

// 示例函数，演示信号量的使用
func worker(ctx *SemaphoreContext, id int) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("Worker %d finished work\n", id)
	case <-ctx.Done():
		fmt.Printf("Worker %d canceled: %v\n", id, ctx.Err())
	}
}

func main() {
	var t uint
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Err(err).Msg("failed to scan input")
		return
	}
	parent := context.Background()
	semaphoreCtx, cancel := NewSemaphoreContext(parent, 3*time.Second)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(semaphoreCtx, id)
		}(i)
	}

	if cancel != nil {
		defer cancel()
	}
	time.Sleep(time.Duration(t) * time.Second)
	// 使用自定义错误取消所有 goroutine
	semaphoreCtx.CancelWithError(errors.New("custom error: semaphore canceled"))

	wg.Wait()
	fmt.Println("All workers done")
}
