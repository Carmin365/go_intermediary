package main

import (
	"errors"
	"fmt"
)

func main() {
	failure := errors.New("An failure occurred")
	if failure != nil {
		fmt.Println(failure)
	}
}
