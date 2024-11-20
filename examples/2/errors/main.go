package main

import (
	"errors"
	"fmt"
)

type zeroDivisionError struct {
	a int
}

func (de *zeroDivisionError) Error() string {
	return fmt.Sprintf("attempted to divide %d by zero", de.a)
}

type nonZeroReminderDivisonError struct {
	a, b int
}

func (de *nonZeroReminderDivisonError) Error() string {
	return fmt.Sprintf("can't divide %d by %d", de.a, de.b)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &zeroDivisionError{a: a}
	}
	if a%b != 0 {
		return 0, &nonZeroReminderDivisonError{a: a, b: b}
	}
	return a / b, nil
}

func DivideAndIncrement(a, b int) (int, error) {
	result, err := Divide(a, b)
	if err != nil {
		return 0, fmt.Errorf("divide: %w", err)
	}
	return result + 1, nil
}

func DisplayDivision(a, b int) {
	result, err := DivideAndIncrement(a, b)
	if err != nil {
		target := &zeroDivisionError{}
		if errors.As(err, &target) {
			fmt.Printf("ZERO DIVISION ERROR: %s\n", err.Error())
		} else {
			fmt.Printf("UNKNOWN ERROR: %s\n", err.Error())
		}
	} else {
		fmt.Println(result)
	}
}

func main() {
	DisplayDivision(10, 5)
	DisplayDivision(10, 0)
	DisplayDivision(10, 3)
}
