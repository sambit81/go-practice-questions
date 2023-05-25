/*
Store 20 numbers in an array then print the maximum and minimum number in the array and also with their index positions.

Example:

# INPUT

If the input array contains numbers like 546
654
851
687
685
7821
211
21325
65464
65487
91654
13348
79825
11645
46546
61684
64678
13454
14564
54654

# OUTPUT

Maximum number: 91654, Index position: 10
Minimum number: 211, Index position: 6
Total numbers found: 6
*/
package main

import "fmt"

const size int = 20 // Constant representing the size of the array

func main() {
	// Declare an array of size 20
	var array [size]int

	fmt.Println("Enter 20 numbers:")

	// Input 20 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// Initialize max and min varible to first element of the array
	max := array[0]
	min := array[0]

	// Initialize the indices to 0
	maxIndex := 0
	minIndex := 0

	// Traverse the array to find maximum and minimum
	for i := 0; i < size; i++ {

		// If the current number is greater than the greatest element then update maximum and its index
		if array[i] > max {
			max = array[i]
			maxIndex = i
		} else if array[i] < min {
			// If the current number is smaller than the smallest element then update minimum and its index
			min = array[i]
			minIndex = i
		}
	}

	// Print the maximum and its index
	fmt.Printf("Maximum number is %d and its index position is %d", max, maxIndex)

	// Print the minimum and its index
	fmt.Printf("Minimum number is %d and its index position is %d", min, minIndex)
}
