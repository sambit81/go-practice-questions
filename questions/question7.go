/*
Store 15 numbers in an array, then sort the numbers present in the array in ascending order.
Sorting algorithm to be used: Selection sort.

Example:

# INPUT

54, 21, 65, 87, 96, 12, 456, 32, 457, 897, 5, 21456, 45, 321, 45

# OUTPUT

5 12 21 32 45 45 54 65 87 96 321 456 457 897 21456

NOTE:

=> In selection sorting, each cell in the array will be compared with rest of the cells.

	For example, a[0] will be compared with a[1] to a[n-1]
				 a[1] will be compared with a[2] to a[n-1] and so on.

=> The outer loop selects the cell and inner loop compares the selected cell with other cells.
=> The selection of the cell will be continued up to a[n-2] while comparison will always be from

	[selected_cell+1] to [n-1] cell.
*/
package main

import "fmt"

func main() {
	// declare an array with length 15
	var arr [15]int

	fmt.Println("Enter 15 numbers")
	// read input integers and store them in the array
	for i := 0; i < 15; i++ {
		fmt.Scan(&arr[i])
	}

	// selection sort algorithm to sort the array in ascending order
	for i := 0; i < 14; i++ {
		for j := i + 1; j < 15; j++ {
			if arr[i] > arr[j] {
				// swap elements
				c := arr[i]
				arr[i] = arr[j]
				arr[j] = c
			}
		}
	}

	// print sorted array in ascending order
	fmt.Println("Numbers in ascending order")
	for i := 0; i < 15; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
