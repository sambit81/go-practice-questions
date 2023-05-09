/*
Write a program that stores 50 integers in an array and then prompts the user to enter
two values: a minimum value and a maximum value. The program should then count and display the frequency
of the numbers that fall within the specified range (inclusive of the minimum and maximum values).

Example:

INPUT
1,2,3,4,5,20,21,22,23,24

Upper bound: 10
Lower bound: 1

OUTPUT
5 [1,2,3,4,5]
*/
package main

import "fmt"

func main() {
	// create an array to store 50 integers
	var arr = [10]int{}

	//get input from user and store in the array
	fmt.Println("Enter 50 numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}

	// input the upper and lower bound
	upper := 0
	lower := 0
	fmt.Print("Enter the upper bound of the range: ")
	fmt.Scan(&upper)

	fmt.Print("Enter the lower bound of the range: ")
	fmt.Scan(&lower)

	// iterate through the array and find elements within the range
	var cnt = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= lower && arr[i] <= upper {
			cnt++
		}
	}

	//display the count of elements
	fmt.Println("The total number of elements found within the range:", cnt)
}
