/*
Store first 20 primes numbers in an array then separate out two-digit prime numbers into a new array. Then, print that new array in reverse order.
Example:

# INPUT

2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71

# OUTPUT

Two digit primes in reverse order:
71 67 61 59 53 47 43 41 37 31 29 23 19 17 13 11
*/
package main

import "fmt"

const size int = 20 // Constant representing the size of the array

func main() {
	// Initialize an array to store the first 20 prime numbers
	var arr [size]int

	// Loop through numbers starting from 2 until 20 primes are found
	for i, j := 0, 2; i < size; j++ {
		num := j
		factors := 0

		// Check the number of factors for the current number
		for k := 1; k <= num; k++ {
			if num%k == 0 {
				factors++
			}
		}

		// If the number has exactly two factors (1 and itself), it is a prime number
		if factors == 2 {
			arr[i] = num
			i++
		}
	}

	// Initialize a slice to store the two-digit prime numbers
	var new_arr []int

	// Loop through the array of primes and add any two-digit prime numbers to the new slice
	for i := 0; i < size; i++ {
		if arr[i] >= 10 && arr[i] <= 99 {
			new_arr = append(new_arr, arr[i])
		}
	}

	// Print the two-digit prime numbers in reverse order
	fmt.Println("Two digits primes in reverse order:")
	for i := len(new_arr) - 1; i >= 0; i-- {
		fmt.Println(new_arr[i])
	}
}
