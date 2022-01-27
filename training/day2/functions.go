package main

import "fmt"

func main() {
	fmt.Println(add(5, 7))
}

func add(a int, b int) int {
	return a + b
}