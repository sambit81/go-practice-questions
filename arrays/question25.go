/*
Store 15 numbers in an array and store number of digits of each number in another array. Sort the digits array in descending order, then display the numbers in
one column and number of digits in another column

Example:

# INPUT

123
45
6789
9876543
987
0
87654
321
54321
12345
987654321
1
9999
222
555555

# OUTPUT

NUMBERS    NUMBER OF DIGITS
987654321       9
9876543         7
555555          6
12345           5
87654           5
54321           5
9999            4
6789            4
321             3
987             3
222             3
123             3
45              2
1               1
0               0
*/
package main

import (
	"fmt"
)

const size int = 15 // Constant representing the size of the array

func main() {
	// Declare number array
	var nums [size]int

	// Declare digits array
	var digits [size]int

	fmt.Println("Enter numbers:")

	// Input numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
	}

	// Generate number of digits for each number
	for i := 0; i < size; i++ {
		// Variable to store duplicate of every number
		dup := nums[i]

		// Counter variable to store count of digits of each number
		count := 0
		for ; dup != 0; dup /= 10 {
			count++
		}
		// Store number of digits
		digits[i] = count
	}

	// Sort the numbers and digit array according to number of digits
	for i := 0; i < size; i++ {
		// Compare each index with all the indices starting from the next index of current index
		for j := i + 1; j < size; j++ {
			if digits[i] < digits[j] {
				// If current index value is smaller than its next index then swap
				d := digits[i]
				digits[i] = digits[j]
				digits[j] = d

				// Numbers has to be swaped if number of digits are swaped
				s := nums[i]
				nums[i] = nums[j]
				nums[j] = s
			}
		}
	}

	// After sorting, digits array will be in descending order and also the numbers array will be sorted accordingly
	// So, print numbers and digits in two columns
	fmt.Println("NUMBERS\t\tNUMBER OF DIGITS")
	for i := 0; i < size; i++ {
		fmt.Printf("%d\t\t%d\n", nums[i], digits[i])
	}
}
