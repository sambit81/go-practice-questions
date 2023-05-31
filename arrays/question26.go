/*
Store 20 numbers in an array. Take a number as input then check how many times the number is present in the array.

Example:

# INPUT

1, 2, 3, 4, 5, 1, 6, 7, 8, 9, 10, 1, 11, 12, 13, 14, 1, 15, 16, 17
Enter a number to check: 1

# OUTPUT

1 is present 4 times in the array
*/
package main

import (
	"fmt"
)

const size int = 20 // Constant representing the size of the array

func main() {
	// Declare number array
	var array [size]int

	fmt.Println("Enter 20 numbers:")

	// Input numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// Variable to store a number which is to be checked
	var num int
	fmt.Print("Enter a number to be checked:")
	fmt.Scan(&num)

	// Counter varible to count the frequency
	cnt := 0

	// Check the frequency
	for i := 0; i < size; i++ {
		// If found, increase the frequency
		if array[i] == num {
			cnt++
		}
	}

	// Display accordingly
	if cnt == 0 {
		fmt.Println("Number not present in the array")
	} else {
		fmt.Printf("Number of times number present: %d\n", cnt)
	}
}
