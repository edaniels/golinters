package main

import "fmt"

func main() {
	// Should be caught by println linter
	fmt.Println("hello world")

	// Should be caught by printf linter
	fmt.Printf("hello %s\n", "world")
}
