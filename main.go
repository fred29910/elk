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
func NewSemaphoreContext(parent context.Context) (*SemaphoreContext, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
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

// Err 实现 context.Context 的 Err 方法，返回自定义错误
func (s *SemaphoreContext) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err != nil {
		return s.err
	}
	return s.Context.Err()
}

// 示例函数，演示信号量的使用
func worker(ctx *SemaphoreContext, id int) {
	select {
	case <-time.After(2 * time.Second):
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
	semaphoreCtx, cancel := NewSemaphoreContext(parent)

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
