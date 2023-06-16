/*
Input exactly 10 numbers, each consisting of two digits only, which will be stored in an array dedicated to two-digit numbers. Similarly, input another set of 10 numbers, each consisting of three digits only, which will be stored in a separate array for three-digit numbers. Once the arrays are populated, find the maximum value among the two-digit numbers and the minimum value among the three-digit numbers.

Example:

Enter only two-digit numbers:
Enter number(1): 45
Enter number(2): 78
Enter number(3): 23
Enter number(4): 566
Number entered is not valid.
Enter number(4): 56
Enter number(5): 12
Enter number(6): 344
Number entered is not valid.
Enter number(6): 34
Enter number(7): 67
Enter number(8): 89
Enter number(9): 76
Enter number(10): 32
Enter only three-digit numbers:
Enter number(1): 234
Enter number(2): 567
Enter number(3): 789
Enter number(4): 7899
Number entered is not valid.
Enter number(4): 123
Enter number(5): 456
Enter number(6): 67
Number entered is not valid.
Enter number(6): 678
Enter number(7): 901
Enter number(8): 354
Enter number(9): 678
Enter number(10): 890

Maximum two-digit value is 89
Minimum three-digit value is 123
*/

package main

import (
	"fmt"
	"math"
)

const size = 10 // Declare a constant for the size of the array

func main() {
	var twoDigit [size]int   // Array to store two-digit numbers
	var threeDigit [size]int // Array to store three-digit numbers

	fmt.Println("Enter only two-digit numbers:")

	cnt := 0  // Counter for the number of two-digit inputs
	temp := 0 // Temporary variable to hold user input
	for cnt < size {
		fmt.Printf("Enter number(%d): ", cnt+1)
		fmt.Scan(&temp)
		// Check if it is a two-digit number
		if temp >= 10 && temp <= 99 {
			twoDigit[cnt] = temp
			cnt++
		} else {
			fmt.Println("Number entered is not valid.")
		}
	}

	fmt.Println("Enter only three-digit numbers:")

	cnt = 0  // Reset the counter for three-digit inputs
	temp = 0 // Reset the temporary variable
	for cnt < size {
		fmt.Printf("Enter number(%d): ", cnt+1)
		fmt.Scan(&temp)
		// Check if it is a three-digit number
		if temp >= 100 && temp <= 999 {
			threeDigit[cnt] = temp
			cnt++
		} else {
			fmt.Println("Number entered is not valid.")
		}
	}

	// Find maximum of two-digit number
	maxTwoDigit := math.MinInt64
	for i := 0; i < size; i++ {
		if twoDigit[i] > maxTwoDigit {
			maxTwoDigit = twoDigit[i]
		}
	}

	// Find minimum of three-digit number
	minThreeDigit := math.MaxInt64
	for i := 0; i < size; i++ {
		if threeDigit[i] < minThreeDigit {
			minThreeDigit = threeDigit[i]
		}
	}

	// Display the results
	fmt.Println("Maximum two-digit value is", maxTwoDigit)
	fmt.Println("Minimum three-digit value is", minThreeDigit)
}
