/*
*
To display 10 numbers stored in an array and their index position in two adjacent
columns with appropriate headings.

Example:

Numbers         Index Position
2                0
5                1
9                2
5                3
10               4
9                5
87               6
56               7
12               8
63               9
*/
package main

import "fmt"

func main() {
	// Declare an array of 10 integers.
	var arr [10]int

	// Prompt the user to enter 10 numbers and store them in the array.
	fmt.Println("Enter 10 numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	// Print the numbers and their index positions in the array.
	fmt.Println("Numbers\t\tIndex Position")
	for index, val := range arr {
		fmt.Println(val, "\t\t", index)
	}
}
