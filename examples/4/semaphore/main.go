package main

import "sync"

type Semaphore struct {
	// TODO: ...
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		// TODO: ...
	}
}

func (s *Semaphore) Acquire() {
	// TODO: ...
}

func (s *Semaphore) Release() {
	// TODO: ...
}

func work(s *Semaphore) {
	s.Acquire()
	defer s.Release()

	// do something
}

func main() {
	s := NewSemaphore(1)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(s)
		}()
	}
	wg.Wait()
}
