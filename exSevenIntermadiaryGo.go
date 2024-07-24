package main

import (
	"fmt"
	"sync"
)

func main() {
	var wG sync.WaitGroup
	wG.Add(1)

	go func() {
		defer wG.Done()
		fmt.Println("Go routine completed")
	}()

	wG.Wait()
	fmt.Println("All go routines concluded")
}
