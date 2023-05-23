/*
Store the first 15 Fibonacci numbers in an array and then display two-digit Fibonacci numbers in descending order.

Example:

# INPUT

If the input array contains numbers like [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377]

# OUTPUT

Two digit Fibonacci numbers in descending order: 89 55 34 21 13
*/
package main

import "fmt"

const size int = 15 // Constant representing the size of the array

func main() {
	// Declare an array of size 15
	var array [size]int

	// Store first two Fibonacci numbers
	array[0] = 0
	array[1] = 1

	// Generate and store first 15 Fibonacci numbers in the array
	for i := 2; i < size; i++ {
		// Every number is sum of its previous two numbers
		array[i] = array[i-1] + array[i-2]
	}

	// Declare a slice to store such numbers
	var two_digit_array = []int{}

	// Traverse the array to find such numbers
	for i := 0; i < size; i++ {

		// If a number is two-digit then add it to the slice 'two_digit_array'
		if array[i] >= 10 && array[i] <= 99 {
			two_digit_array = append(two_digit_array, array[i])
		} else if array[i] >= 100 {
			// If at any point we encounter a three-digit number then break the loop.
			// After this point there will be no two-digit number
			break
		}
	}

	// Print the slice in reverse order
	fmt.Println("Two digit Fibonacci numbers in descending order:")
	for i := len(two_digit_array) - 1; i >= 0; i-- {
		fmt.Printf("%d ", two_digit_array[i])
	}

}
