/*
Create an array to store 30 numbers. Subsequently, identify the positions of the maximum and minimum numbers within the array. Finally, calculate the sum of the numbers located between the maximum and minimum numbers, inclusive of both the maximum and minimum values.

Example:

# INPUT
5
6
14
98
78
32
15
68
97
1

# OUTPUT
Sum of numbers which lies between minimum number in array: 1 and maximum number in array: 98 is 389
*/

package main

import (
	"fmt"
	"math"
)

const size int = 30 // Declare a constant for the size of the array

func main() {

	var num [size]int

	maxIndex := 0 // Initialize variable to store the index of the maximum number
	minIndex := 0 // Initialize variable to store the index of the minimum number

	fmt.Println("Please enter", size, "numbers:")

	// Read numbers from input and store them in the array
	for i := 0; i < size; i++ {
		fmt.Scan(&num[i])
	}

	// Find the index of the maximum and minimum numbers in the array
	for i := 0; i < size; i++ {
		if num[i] > num[maxIndex] {
			maxIndex = i
		}

		if num[i] < num[minIndex] {
			minIndex = i
		}
	}

	// Determine the start and end index of the range of numbers between the maximum and minimum numbers
	startIndex := int(math.Min(float64(maxIndex), float64(minIndex)))
	endIndex := int(math.Max(float64(maxIndex), float64(minIndex)))

	sum := 0 // Initialize variable to store the sum of numbers in the range

	// Calculate the sum of numbers in the range
	for i := startIndex; i <= endIndex; i++ {
		sum += num[i]
	}

	// Display the results
	fmt.Printf("The position of the maximum number in the array is: %d\n", num[maxIndex])
	fmt.Printf("The position of the minimum number in the array is: %d\n", num[minIndex])
	fmt.Printf("The sum of the numbers between the minimum number %d and the maximum number %d (inclusive) is: %d\n", num[minIndex], num[maxIndex], sum)

}
