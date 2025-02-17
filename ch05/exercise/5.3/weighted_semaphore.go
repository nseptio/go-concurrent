package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *Semaphore) Acquire(permits int) {
	rw.cond.L.Lock()
	for rw.permits-permits < 0 {
		rw.cond.Wait()
	}
	rw.permits -= permits
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release(permits int) {
	rw.cond.L.Lock()
	rw.permits += permits
	rw.cond.Signal()
	rw.cond.L.Unlock()
}

func main() {
	sema := NewSemaphore(3)
	sema.Acquire(2)
	fmt.Println("Parent thread acquired semaphore")
	go func() {
		sema.Acquire(2)
		fmt.Println("Child thread acquired semaphore")
		sema.Release(2)
		fmt.Println("Child thread released semaphore")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Parent thread releasing semaphore")
	sema.Release(2)
	time.Sleep(1 * time.Second)
}
