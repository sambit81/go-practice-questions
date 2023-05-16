/*
Store 10 numbers in an array and then display the number in one column and whether it is positive or negative in another column.

Example:

# INPUT

5
97
-8
-3
2
5
8
1
6
-4

# OUTPUT

5               Positive
97              Positive
-8              Negative
-3              Negative
2               Positive
5               Positive
8               Positive
1               Positive
6               Positive
-4              Negative
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

	fmt.Println("Numbers\t\tSign")

	// Check whether each number is +ve or -ve
	for i := 0; i < size; i++ {
		if array[i] >= 0 {
			fmt.Printf("%d\t\tPositive\n", array[i])
		} else {
			fmt.Printf("%d\t\tNegative\n", array[i])
		}
	}
}
