package main

import "fmt"

func main() {
	var v any = "asd"

	switch t := v.(type) {
	case int:
		fmt.Println("got int", t)

	case string:
		fmt.Println("got string", t)

	default:
		fmt.Println("unknown type")
	}
}
