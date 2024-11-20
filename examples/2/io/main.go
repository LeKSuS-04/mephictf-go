package main

import (
	"fmt"
	"io"
	"os"
)

func ReadAndReverse(r io.Reader) {
	b := make([]byte, 4096)
	n, err := r.Read(b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	read := b[:n]
	for i := 0; i < n/2; i++ {
		read[i], read[n-i-1] = read[n-i-1], read[i]
	}
	fmt.Println(string(read))
}

func main() {
	ReadAndReverse(os.Stdin)
}
