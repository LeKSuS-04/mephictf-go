package main

import "fmt"

func sum(nums []int) int {
	nums[0] = 0
	nums = append(nums, 5)
	nums = append(nums, 5)
	fmt.Println(nums)
	nums[0] = 8
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func main() {
	v := [4]int{1, 2, 3, 4}
	// sli	 |-|
	// arr 1 2 3 4
	// ptr   ^

	// sli	 |---|
	// arr 1 2 3 5
	// ptr   ^
	fmt.Println(sum(v[1:3]))
	fmt.Println(v)
}
