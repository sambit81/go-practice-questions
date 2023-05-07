/*
Create an array of 15 nos.
Ask the user for input for 15 integers. Display only the even numbers present in the array and also display the count of it.

Example:

INPUT
2
5
8
6
1
4
6
8
2
3
6
98
20
21
25

OUTPUT
2
8
6
4
6
8
2
6
98
20
Total number of even(s): 10
*/
package main

import "fmt"

func main() {
	// create an array of 15 integers
	var arr = [15]int{}

	// get input from user and store in the array
	fmt.Println("Enter 15 numbers:")
	for i := 0; i < 15; i++ {
		fmt.Scan(&arr[i])
	}

	// iterate through the array and find even numbers
	var cnt = 0
	fmt.Println("The even numbers present in the array are: ")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Println(arr[i])
			cnt++
		}
	}

	// display the count of even numbers
	fmt.Println("The total number of even(s) present in array:", cnt)
}
