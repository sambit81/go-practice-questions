/*
Store 20 nos in an array. Display all zeroes first (if any) followed by all the numbers without changing their sequence.

Example:

INPUT
23
44
5
-3
0
43
67
8
0
29
30
0
21
40
14
51
99
0
19
-88

OUTPUT
0 0 0 0 23 44 5 -3 43 67 8 29 30 21 40 14 51 99 19 -88
*/

package main

import "fmt"

func main() {
	// Initializing an array of size 20 with 0
	var arr = [20]int{}

	// Taking 20 inputs from user
	fmt.Println("Enter 20 numbers")
	for i := 0; i < 20; i++ {
		fmt.Scan(&arr[i])
	}

	// Printing all the zeros present in the array
	for i := 0; i < 20; i++ {
		if arr[i] == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// Printing all the non-zero values present in the array
	for i := 0; i < 20; i++ {
		if arr[i] != 0 {
			fmt.Print(arr[i], " ")
		}
	}
}
