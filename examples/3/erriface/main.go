package main

import "fmt"

type ZeroDivisionError struct {
}

func (*ZeroDivisionError) Error() string {
	return "zero division error"
}

func CheckDivision(a, b int) *ZeroDivisionError {
	if b == 0 {
		return &ZeroDivisionError{}
	}
	return nil
}

func DoDivision(a, b int) (int, error) {
	var err error
	if err = CheckDivision(a, b); err != nil {
		return 0, err
	}
	return a / b, nil
}

func main() {
	n, err := DoDivision(10, 5)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	} else {
		fmt.Println("Got", n)
	}
}
