/*
Store 35 numbers in an array and then display how many prime numbers are in the range (25-50) and how many primes are not in the range.

Example:

# INPUT

If the input array contains numbers like [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]
Range : [5 - 10]

Primes in the range - 2 (5,7)

# OUTPUT

Primes not in range - 4 (6,8,9,10)
*/
package main

import "fmt"

const size int = 35 // Constant representing the size of the array

func main() {
	// Declare an array of size 35
	var array [size]int

	// Lower bound of the range
	lower := 25

	// Upper bound of the range
	upper := 50

	fmt.Println("Enter 35 numbers:")

	// Input 35 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	// Initialize a counter for numbers in range
	rn := 0

	// Initialize a counter for numbers not in range
	nt_rn := 0

	for i := 0; i < size; i++ {

		// If the number is within the range, then check for prime or not
		if array[i] >= lower && array[i] <= upper {

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
				rn++
			}
		} else {
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
				rn++
			}
		}
	}

	// Print number of primes within the range and not within the range
	fmt.Printf("Prime numbers in range: %d\n", rn)
	fmt.Printf("Prime numbers not in range: %d", nt_rn)

}
