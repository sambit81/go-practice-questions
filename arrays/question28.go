/*
Take a five digit number as input, and store all its digit in an array. Use any sorting algorithm to sort the digits array in descending order. At last, display the sorted digits array.

Example:

# INPUT

45678

# OUTPUT

[8 7 6 5 4]
*/

package main

import (
	"fmt"
)

func main() {
	// Declare a variable to store the five digit number
	var num int

	// Read input from the user
	fmt.Print("Enter a five-digit number: ")
	fmt.Scan(&num)

	// Declare an array to store the digits
	var digit [5]int

	// Extracting digits from the number and storing them in the array
	dup := num
	index := 1
	for ; dup != 0; dup /= 10 {
		digit[5-index] = dup % 10
		index++
	}

	// Sorting the digits array in descending order using Selection Sort
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			// Swap elements if the current element is less than the next element
			if digit[i] < digit[j] {
				d := digit[i]
				digit[i] = digit[j]
				digit[j] = d
			}
		}
	}

	// Display the sorted digits array
	fmt.Println("Digits in descending order: ", digit)
}
