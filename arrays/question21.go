/*
Store the first 15 Fibonacci numbers in an array and then display only prime Fibonacci numbers in descending order and also display how many such numbers found.

Example:

# INPUT

If the input array contains numbers like [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377]

# OUTPUT

Prime Fibonacci numbers in descending order: 233 89 13 5 3 2
Total numbers found: 6
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
	var prime_array = []int{}

	// Traverse the array to find such numbers
	for i := 0; i < size; i++ {
		// Initialize a counter for factors
		cnt := 0

		// Loop from 1 to the current number and check for factors
		for j := 1; j <= array[i]; j++ {
			// If found a factor, then increase the factor counter
			if array[i]%j == 0 {
				cnt++
			}
		}

		// Check if number of factors is 2 for prime
		if cnt == 2 {
			prime_array = append(prime_array, array[i])
		}
	}

	// Print the slice in reverse order
	fmt.Println("Prime Fibonacci numbers in descending order:")
	for i := len(prime_array) - 1; i >= 0; i-- {
		fmt.Printf("%d ", prime_array[i])
	}
	fmt.Printf("\nTotal such numbers found: %d", len(prime_array))
}
