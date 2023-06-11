/*
Store 10 numbers in an array. Then store the reverse of each number in another array. Print only those pairs whose difference is below 100.

Example:

# INPUT

45678

# OUTPUT

[8 7 6 5 4]
*/

package main

import (
	"fmt"
	"math"
)

const size int = 10

func main() {
	// Declare arrays
	var num [10]int
	var rev_num [10]int

	fmt.Println("Enter 10 numbers:")
	for i := 0; i < size; i++ {
		fmt.Scan(&num[i]) // Read user input and store in num array
	}

	for i := 0; i < size; i++ {
		num_len := -1
		dup := num[i]
		for ; dup != 0; dup /= 10 {
			num_len++ // Calculate the number of digits in the number
		}
		dup = num[i]
		sum := 0
		for ; dup != 0; dup /= 10 {
			digit := dup % 10
			sum += digit * int(math.Pow(float64(10), float64(num_len))) // Reverse the number
			num_len--
		}
		rev_num[i] = sum // Store the reversed number in rev_num array
	}

	fmt.Println("Pairs:")
	for i := 0; i < size; i++ {
		if int(math.Abs(float64(num[i]-rev_num[i]))) < 100 {
			fmt.Printf("%d\t%d\n", num[i], rev_num[i]) // Print pairs whose difference is below 100
		}
	}
}
