package main

import (
	"fmt"
	"ioutil"
)

func main() {
	content := "Hello, Go!"

	err := ioutil.WriteFile("test.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Println("File created successfully!")
}
