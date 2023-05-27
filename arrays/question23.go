/*
Create two arrays and store 15 numbers in each of the array. Then create another array to store the difference of elements of both the arrays.
Then print all the three arrays where the difference is more than 100.

Example:

# INPUT

First array:
110
180
220
350
450
500
650
750
820
910
980
1050
1120
1500
2000

Second array:
80
120
150
280
400
420
530
620
710
780
850
940
1020
2500
3000

# OUTPUT
First  Second  Difference
650     220     120
750     220     130
820     220     110
910     220     130
980     220     130
1050    220     110
1500    220     1000
2000    220     1000
*/
package main

import (
	"fmt"
	"math"
)

const size int = 15 // Constant representing the size of the array

func main() {
	// Declare first array of size 15
	var array1 [size]int

	// Declare second array of size 15
	var array2 [size]int

	fmt.Println("Enter numbers for first array")

	// Input 15 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array1[i])
	}

	fmt.Println("Enter numbers for second array")

	// Input 15 numbers
	for i := 0; i < size; i++ {
		fmt.Scan(&array2[i])
	}

	// Declare another array to store the difference of the elements of above two arrays
	var difference [size]int

	// Store the absolute difference between elements in the 'difference' array
	for i := 0; i < size; i++ {
		difference[i] = int(math.Abs(float64(array1[i]) - float64(array2[i])))
	}

	// Prints elements from all three arrays where difference is more than 100
	fmt.Println("First\tSecond\tDifference")
	for i := 0; i < size; i++ {
		if difference[i] > 100 {
			fmt.Printf("%d\t%d\t%d\n", array1[i], array1[2], difference[i])
		}
	}
}
