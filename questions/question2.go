/*
Store and display 10 numbers in one column and their square in another column.

Example:

Numbers         Square
1                1
2                4
3                9
4                16
5                25
6                36
7                49
8                64
9                81
10               100
*/
package main

import "fmt"

func main() {
	// Declare an array of 10 integers
	var arr [10]int

	// Prompt the user to enter 10 numbers and store them in the array
	fmt.Println("Enter 10 numbers")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	// Print the entered numbers and their squares
	fmt.Println("Numbers\t\tSquare")
	for _, val := range arr {
		fmt.Println(val, "\t\t", val*val)
	}
}
