package main

import (
	"fmt"
	"iter"
)

func Counter() iter.Seq[int] {
	return func(yield func(int) bool) {
		run := true
		for i := 0; run; i++ {
			run = yield(i)
		}
	}
}

func Double(it iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for value := range it {
			x := value * 2
			if !yield(x) {
				return
			}
		}
	}
}

func main() {
	for v := range Double(Counter()) {
		fmt.Println(v)
	}
}
