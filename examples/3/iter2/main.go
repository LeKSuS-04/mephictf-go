package main

import (
	"fmt"
	"iter"
	"slices"
)

func Enumerate(it iter.Seq[int]) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		i := 0
		for v := range it {
			if !yield(i, v) {
				return
			}
			i += 1
		}
	}
}

func main() {
	numbers := []int{3, 7, 9, 42, 2, 8, 13}

	for k, v := range Enumerate(slices.Values(numbers)) {
		fmt.Printf("%d: %d\n", k, v)
	}
}
