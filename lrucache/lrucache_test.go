package lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCacheEmpty(t *testing.T) {
	t.Parallel()
	c := New(0)

	c.Set(1, 2)

	_, ok := c.Get(1)
	require.False(t, ok)
}

func TestCacheUpdate(t *testing.T) {
	t.Parallel()
	c := New(1)

	_, ok := c.Get(1)
	require.False(t, ok)

	c.Set(1, 2)

	v, ok := c.Get(1)
	require.True(t, ok)
	require.Equal(t, 2, v)
}

func TestCacheRange(t *testing.T) {
	t.Parallel()
	c := New(5)

	for i := 0; i < 5; i++ {
		c.Set(i, i+1)
	}

	c.Get(0)
	c.Get(1)
	c.Set(5, 6)
	c.Set(6, 7)

	var keys, values []int

	c.Range(func(k, v int) bool {
		keys = append(keys, k)
		values = append(values, v)
		return true
	})
	require.Equal(t, []int{4, 0, 1, 5, 6}, keys)
	require.Equal(t, []int{5, 1, 2, 6, 7}, values)
}

func TestCacheClear(t *testing.T) {
	t.Parallel()
	c := New(5)

	for i := 0; i < 5; i++ {
		c.Set(i, i)
	}

	for i := 0; i < 5; i++ {
		v, ok := c.Get(i)
		require.True(t, ok)
		require.Equal(t, i, v)
	}

	c.Clear()

	for i := 0; i < 5; i++ {
		_, ok := c.Get(i)
		require.False(t, ok)
	}
}

func TestCacheRangeWithStop(t *testing.T) {
	t.Parallel()
	c := New(5)

	for i := 0; i < 5; i++ {
		c.Set(i, i)
	}

	var keys, values []int
	c.Range(func(k, v int) bool {
		return false
	})
	require.Empty(t, keys)
	require.Empty(t, values)

	c.Range(func(k, v int) bool {
		keys = append(keys, k)
		values = append(values, v)
		return k < 3
	})
	require.Equal(t, []int{0, 1, 2, 3}, keys)
	require.Equal(t, []int{0, 1, 2, 3}, values)
}
