package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var sli1 []int = nil
	var sli2 []int = []int{}
	var sli3 []int = make([]int, 0)

	s1 := (*reflect.SliceHeader)(unsafe.Pointer(&sli1))
	s2 := (*reflect.SliceHeader)(unsafe.Pointer(&sli2))
	s3 := (*reflect.SliceHeader)(unsafe.Pointer(&sli3))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
