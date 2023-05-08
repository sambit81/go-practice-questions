/*
Store 15 numbers in an array, then sort the numbers present in the array in ascending order.
Sorting algorithm to be used: Bubble sort.

Example:

# INPUT

54, 98 ,2 ,54 ,96 ,78 ,32 ,59 ,78 ,12 ,534 ,979 ,543 ,515 ,87

# OUTPUT

2 12 32 54 54 59 78 78 87 96 98 515 534 543 979

NOTE:

=> In Bubble Sorting, each time, two consecutive cells in the array will be compared.

	For example, a[0] will be compared with a[1]
				 a[1] will be compared with a[2] and so on.

=> After the end of the checking, the last cell will carry the highest number in case of ascending order

	or lowest number in case of descending order.

=> The outer loop helps to remove cells from the last part of the array.
=> The selection of the cell will be continued up to a[n-2] while comparison will always be from

	[selected_cell+1] to [n-1] cell.
*/
package main

import "fmt"

func main() {
	// declare an array of 15 integers
	var arr = [15]int{}

	fmt.Println("Enter 15 numbers:")

	// get input from user and store it in the array
	for i := 0; i < 15; i++ {
		fmt.Scan(&arr[i])
	}

	// use bubble sort to sort the array in ascending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			// compare adjacent elements and swap them if they are in the wrong order
			if arr[j] > arr[j+1] {
				c := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = c

				// Alternative to above swap mechanism
				//arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Numbers in ascending order:")

	// print the sorted array
	for i := 0; i < 15; i++ {
		fmt.Println(arr[i])
	}
}
