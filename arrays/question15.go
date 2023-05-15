/*
Store 10 numbers in an array and print only armstrong numbers from the array along with its index.

Armstrong number: A number will be called as armstrong number if it is equal to sum of the cube of its digits.

Example -> 407 (407 = 64 + 0 + 343), 371 (371 = 27 + 343 + 1)

Example:

# INPUT

0
1
371
1000
2500
407
153
8028
8208
9474

# OUTPUT

0 ,index= 0
1 ,index= 1
371 ,index= 2
407 ,index= 5
153 ,index= 6
*/
package main

import "fmt"

const size int = 10 // Constant representing the size of the array

func main() {
	var array [size]int

	fmt.Println("Enter 10 numbers:")

	// Input
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Armstrong numbers found:")

	// Check for Armstrong numbers
	for i := 0; i < size; i++ {
		num := array[i]
		sum := 0

		// Calculate the sum of cubes of digits
		for ; num != 0; num /= 10 {
			digit := num % 10
			sum += digit * digit * digit
		}

		// Check if the sum is equal to the original number
		if sum == array[i] {
			fmt.Println(array[i], ",index=", i)
		}
	}
}
