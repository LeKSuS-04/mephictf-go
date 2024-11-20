package main

import "fmt"

type PalindromeDetector interface {
	IsPalindrome() bool
}

type MyStruct struct {
}

type MyString string

func (s MyString) IsPalindrome() bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

type MyIntSlice []int

func (s MyIntSlice) IsPalindrome() bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []any{
		MyString("hello"),
		1,
		2,
		"bebebe",
		MyIntSlice{1, 2, 3, 2, 1},
		MyIntSlice{1, 2, 3, 2, 4},
		MyString("olololo"),
		map[int]int{1: 2, 3: 4},
		MyString("olololo"),
		MyString("olala"),
		MyStruct{},
		nil,
	}

	for _, item := range arr {
		if p, ok := item.(PalindromeDetector); ok {
			fmt.Println(p)
		}
	}
}
