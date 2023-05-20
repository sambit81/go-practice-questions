/*
Store 13 numbers in an array and then display the positive numbers first followed by the negative numbers without altering the sequence.

Example:

# INPUT

2
3
-5
22
-8
5
77
-88
3
-9
45
-23
-2

# OUTPUT

Positive numbers:
2 3 22 5 77 3 45
Negative numbers:
-5 -8 -88 -9 -23 -2
*/
package main

import "fmt"

const size int = 13 // Constant representing the size of the array

func main() {
	// Declare an array of size 13
	var array [size]int

	fmt.Println("Enter 13 numbers:")

	// Input 13 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// Loop to print positive numbers first
	fmt.Println("Positive numbers:")
	for i := 0; i < size; i++ {
		if array[i] >= 0 {
			fmt.Print(array[i], " ")
		}
	}

	// Just to print a new line
	fmt.Println()

	// Loop to print negative numbers
	fmt.Println("Negative numbers:")
	for i := 0; i < size; i++ {
		if array[i] < 0 {
			fmt.Print(array[i], " ")
		}
	}
}
