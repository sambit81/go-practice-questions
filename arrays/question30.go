/*
Create an array which will store 25 two-digit numbers. If the user enters a number which is not two-digit then it should show an error message.
The program should terminate when the user has entered exact 25 two-digit numbers.

Example:

Enter numbers:
Enter number(1): 54
Enter number(2): 65
Enter number(3): 98
Enter number(4): 12
Enter number(5): 111
Number entered is not valid.
Enter number(5): 32
Enter number(6): 12
Enter number(7): 54
Enter number(8): 96
Enter number(9): 369
Number entered is not valid.
Enter number(9): 657
Number entered is not valid.
Enter number(9): 987
Number entered is not valid.
Enter number(9): 12
Enter number(10): 45
Enter number(11): 65
Enter number(12): 78
Enter number(13): 89
Enter number(14): 84
Enter number(15): 36
Enter number(16): 96
Enter number(17): 78
Enter number(18): 98
Enter number(19): 12
Enter number(20): 45
Enter number(21): 65
Enter number(22): 1555
Number entered is not valid.
Enter number(22): 321
Number entered is not valid.
Enter number(22): 212
Number entered is not valid.
Enter number(22): 54
Enter number(23): 65
Enter number(24): 98
Enter number(25): 100
Number entered is not valid.
Enter number(25): 12
*/

package main

import (
	"fmt"
)

const size int = 25 // Declare a constant which will be size of the arrays

func main() {
	// Declare array
	var num [size]int

	// Counter varible to store count of index
	cnt := 0

	// Temp varible to take input
	temp := 0

	fmt.Println("Enter numbers:")
	for true && cnt < 25 {
		fmt.Printf("Enter number(%d): ", cnt+1)
		fmt.Scan(&temp)
		// Check if it is a two-digit number
		if temp >= 10 && temp <= 99 {
			num[cnt] = temp
			cnt++
		} else {
			fmt.Println("Number entered is not valid.")
		}
	}
}
