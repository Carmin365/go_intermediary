package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, annul := context.WithTimeout(context.Background(), 2*time.Second)
	defer annul()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("operation successful")
	case <-ctx.Done():
		fmt.Println("operation timed out")
	}
}
