package main

import (
	"log"
	"sync"
)

const (
	N           = 1e10
	NUM_WORKERS = 4
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func simple() float64 {
	log.Println("Calculating sum")
	sum := 0.0
	for i := 0; i < N; i++ {
		sum += float64(i)
	}
	log.Println("Sum:", sum)
	return sum
}

func parallel() float64 {
	log.Println("Calculating sum")
	sums := make([]float64, NUM_WORKERS)

	wg := sync.WaitGroup{}
	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := i * N / NUM_WORKERS; j < (i+1)*N/NUM_WORKERS; j++ {
				sums[i] += float64(j)
			}
		}(i)
	}
	wg.Wait()

	sum := 0.0
	for _, s := range sums {
		sum += s
	}
	log.Println("Sum:", sum)
	return sum
}

func main() {
	log.Println("Calculating sum")
	sum := simple()
	// sum := parallel()
	log.Println("Sum:", sum)
}
