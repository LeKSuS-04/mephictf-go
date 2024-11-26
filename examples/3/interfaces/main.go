package main

import (
	"fmt"
	"strconv"
)

type Binary uint64

func (b Binary) String() string {
	return strconv.FormatUint(uint64(b), 10)
}

func main() {
	b := Binary(1)
	var iface fmt.Stringer = b
	fmt.Println(iface)
}
