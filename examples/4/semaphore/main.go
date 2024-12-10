package main

import "sync"

type Semaphore struct {
	c chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		c: make(chan struct{}, n),
	}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

func work(s *Semaphore) {
	s.Acquire()
	defer s.Release()

	// do something
}

func main() {
	s := NewSemaphore(3)
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
