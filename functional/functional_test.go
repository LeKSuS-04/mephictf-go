package functional

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStream(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty_slice",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "single_element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple_elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			result := slices.Collect(stream.Iterate())
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		f        func(int) int
		expected []int
	}{
		{
			name:     "double_numbers",
			input:    []int{1, 2, 3},
			f:        func(x int) int { return x * 2 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "empty_stream",
			input:    []int{},
			f:        func(x int) int { return x * 2 },
			expected: nil,
		},
		{
			name:     "add_one",
			input:    []int{0, 1, 2},
			f:        func(x int) int { return x + 1 },
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			mapped := stream.Map(tc.f)
			result := slices.Collect(mapped.Iterate())
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "even_numbers",
			input:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  []int{2, 4, 6},
		},
		{
			name:      "empty_stream",
			input:     []int{},
			predicate: func(x int) bool { return x > 0 },
			expected:  nil,
		},
		{
			name:      "positive_numbers",
			input:     []int{-2, -1, 0, 1, 2},
			predicate: func(x int) bool { return x > 0 },
			expected:  []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			filtered := stream.Filter(tc.predicate)
			result := slices.Collect(filtered.Iterate())
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamTake(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "take_some",
			input:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "take_none",
			input:    []int{1, 2, 3},
			n:        0,
			expected: []int{},
		},
		{
			name:     "take_more_than_length",
			input:    []int{1, 2, 3},
			n:        5,
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			taken := stream.Take(tc.n)
			result := slices.Collect(taken.Iterate())
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamDrop(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "drop_some",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5},
		},
		{
			name:     "drop_none",
			input:    []int{1, 2, 3},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "drop_all",
			input:    []int{1, 2, 3},
			n:        3,
			expected: []int{},
		},
		{
			name:     "drop_more_than_length",
			input:    []int{1, 2, 3},
			n:        5,
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			dropped := stream.Drop(tc.n)
			result := slices.Collect(dropped.Iterate())
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamFoldLeft(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		f        func(int, int) int
		expected int
	}{
		{
			name:     "sum",
			input:    []int{1, 2, 3, 4, 5},
			f:        func(acc, x int) int { return acc + x },
			expected: 15,
		},
		{
			name:     "product",
			input:    []int{1, 2, 3, 4},
			f:        func(acc, x int) int { return acc * x },
			expected: 24,
		},
		{
			name:     "empty_stream",
			input:    []int{},
			f:        func(acc, x int) int { return acc + x },
			expected: 0,
		},
		{
			name:     "left_to_right_order",
			input:    []int{100, 10, 2},
			f:        func(acc, x int) int { return acc - x },
			expected: 88,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			result := stream.FoldLeft(tc.f)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStreamForEach(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "collect_values",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty_stream",
			input:    []int{},
			expected: nil,
		},
		{
			name:     "single_element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stream := NewStream(slices.Values(tc.input))
			var result []int
			stream.ForEach(func(x int) {
				result = append(result, x)
			})
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestInfiniteSequence(t *testing.T) {
	count := func(x int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for {
				if !yield(x) {
					break
				}
				x += 1
			}
		}
	}

	stream := NewStream(count(1))
	result := stream.
		Map(func(x int) int { return x * 2 }).
		Drop(10).
		Filter(func(x int) bool { return x%3 == 0 }).
		Take(5).
		FoldLeft(func(acc, x int) int { return acc + x })

	require.Equal(t, 180, result)
}

func TestStreamCombined(t *testing.T) {
	stream := NewStream(slices.Values([]int{1, 2, 3, 4, 5, 6, 7}))

	it := stream.
		Map(func(x int) int { return x * 3 }).
		Filter(func(x int) bool { return x%2 != 0 }).
		Drop(1).
		Take(2).
		Map(func(x int) int { return x + 1 }).
		Iterate()

	result := slices.Collect(it)
	require.Equal(t, []int{10, 16}, result)
}
