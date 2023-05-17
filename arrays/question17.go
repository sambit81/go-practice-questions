/*
Store 10 numbers in an array and then display the maximum number present in the array and also print its index position.

Example:

# INPUT

157
27
76
89
100
2000
5000
231
54
96

# OUTPUT

Maximum number: 5000
Index position: 6
*/
package main

import "fmt"

const size int = 10 // Constant representing the size of the array

func main() {
	// Declare an array of size 10
	var array [size]int

	fmt.Println("Enter 10 numbers:")

	// Input 10 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// Declare max and index variable to store maximum number and its index position
	max := 0
	index := -1

	// Find out the maximum number in the array
	for i := 0; i < size; i++ {
		if array[i] > max {
			max = array[i]
			index = i
		}
	}

	// Display the maximum along with its index
	fmt.Printf("Maximum number in the array is %d and its index position is: %d", max, index)

}
