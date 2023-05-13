/*
Store 15 numbers in an array and separate out those numbers with 3 digits or more; all the digits are the same. After finding such numbers print those numbers.

Example:

# INPUT

3333
1111
2222
654
548
333
11111
546565
12
3333
4444
5555
1021
3214
5479

# OUTPUT

3333 1111 2222 333 11111 3333 4444 5555
*/
package main

import "fmt"

const size = 15 // constant representing size of array

func main() {
	var arr [size]int // declare an array of size 15 to store the numbers

	fmt.Println("Enter 15 numbers:")
	for i := 0; i < size; i++ { // loop through the array and take input from user
		fmt.Scan(&arr[i])
	}

	var new_arr []int // declare a slice to store the numbers which have 3 digits or more and all digits are same

	for i := 0; i < size; i++ {
		num := arr[i]      // take the current number
		if arr[i] >= 100 { // check if it has 3 digits or more
			ld := num % 10              // store the last digit
			for ; num != 0; num /= 10 { // loop through the digits of the number
				d := num % 10 // take the current digit
				if d != ld {  // check if it's not equal to the last digit
					break // if not, break the loop
				}
			}
			if num == 0 { // if the loop completed successfully, and all digits are same, add it to the slice
				new_arr = append(new_arr, arr[i])
			}
		}
	}

	fmt.Println("Numbers found:", new_arr) // print the numbers which have 3 digits or more and all digits are same
}
