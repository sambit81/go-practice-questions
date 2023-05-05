/*
Store 11 to 19 in reverse order in an array and then display those numbers in ascending order.

Example:

11
12
13
14
15
16
17
18
19
*/
package main

import "fmt"

func main() {
	// Initialize an integer array of length 9
	var arr [9]int

	// Loop through the array and assign values in reverse order
	for i := 0; i < 9; i++ {
		arr[i] = 19 - i
	}

	// Print the array in reverse order
	fmt.Println("Values in reverse order:", arr)

	// Print the array in correct order by iterating from the last index to the first index
	fmt.Println("Values in correct order:")
	for i := 8; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
