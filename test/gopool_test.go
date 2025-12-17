package test

import (
	"sync"
	"testing"
	"time"

	"zecs/zecs"
)

func TestGoPool_RunAndWait(t *testing.T) {
	var mu sync.Mutex
	results := []int{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		val := i
		wg.Add(1)
		zecs.GO(func() {
			time.Sleep(10 * time.Millisecond)
			mu.Lock()
			results = append(results, val)
			mu.Unlock()
			wg.Done()
		})
	}
	wg.Wait()

	if len(results) != 10 {
		t.Errorf("expected 10 results, got %d", len(results))
	}
}

func TestGoPool_ConcurrencyLimit(t *testing.T) {
	var running int
	var mu sync.Mutex
	maxRunning := 0
	var wg sync.WaitGroup
	sem := make(chan struct{}, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		zecs.GO(func() {
			sem <- struct{}{} // acquire
			mu.Lock()
			running++
			if running > maxRunning {
				maxRunning = running
			}
			mu.Unlock()
			time.Sleep(20 * time.Millisecond)
			mu.Lock()
			running--
			mu.Unlock()
			<-sem // release
			wg.Done()
		})
	}
	wg.Wait()

	if maxRunning > 2 {
		t.Errorf("expected max concurrency 2, got %d", maxRunning)
	}
}

func TestGoPool_PanicRecovery(t *testing.T) {
	var called bool
	var wg sync.WaitGroup
	wg.Add(1)
	zecs.GO(func() {
		defer func() {
			if r := recover(); r != nil {
				called = true
			}
			wg.Done()
		}()
		panic("test panic")
	})
	wg.Wait()

	if !called {
		t.Error("expected panic to be recovered")
	}
}

func TestGoPool_WaitWithoutGo(t *testing.T) {
	// 直接测试 WaitGroup 无 goroutine 的情况
	var wg sync.WaitGroup
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("Wait should return immediately when no goroutines")
	}
}
