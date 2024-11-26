package main

import (
	"fmt"
)

func modifyMap(m map[string]int) {
	m["a"] = 10
}

func main() {
	var m map[string]int = make(map[string]int)
	fmt.Println(m)
	modifyMap(m)
	fmt.Println(m)
}
