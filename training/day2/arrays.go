package main

import "fmt"

var numbersArray [5]int
var namesArray [3]string

func main() {
	numbersArray[0] = 10
	numbersArray[1] = 20
	numbersArray[2] = 30
	numbersArray[3] = 40
	numbersArray[4] = 50

	fmt.Println("Numbers array is:", numbersArray)

	namesArray[0] = "Jimmy"
	namesArray[1] = "John"
	namesArray[2] = "Ron"

	fmt.Println("Names array is:", namesArray)
}