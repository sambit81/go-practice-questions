/*
Store 35 numbers in an array and then display how many prime numbers are in the range (25-50) and how many primes are not in the range.

Example:

# INPUT

If the input array contains numbers like [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]
Range : [5 - 10]

Primes in the range - 2 (5,7)

# OUTPUT

Primes not in range - 4 (6,8,9,10~)
*/
package main

import "fmt"

const size int = 35 // Constant representing the size of the array

func main() {
	// Declare an array of size 35
	var array [size]int

	lower := 25
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
		if array[i] >= lower && array[i] <= upper {
			cnt := 0
			for j := 1; j <= array[i]; j++ {
				if array[i]%j == 0 {
					cnt++
				}
			}
			if cnt == 2 {
				rn++
			}
		} else {
			cnt := 0
			for j := 1; j <= array[i]; j++ {
				if array[i]%j == 0 {
					cnt++
				}
			}
			if cnt == 2 {
				nt_rn++
			}
		}
	}

	fmt.Printf("Prime numbers in range: %d\n", rn)
	fmt.Printf("Prime numbers not in range: %d", nt_rn)

}
