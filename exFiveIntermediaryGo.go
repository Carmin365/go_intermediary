package main

import (
	"errors"
	"fmt"
)

func to_divide(dividend int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}
	return dividend / divider, nil
}

func main() {
	number1 := 10
	number2 := 2

	result, err := to_divide(number1, number2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}
}
