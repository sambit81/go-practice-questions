/*
Store marks of 10 students in an array and then print the frequency of marks obtained by students in the range:

Range
=====
0-40
41-60
61-80
81-100

Example:

# INPUT

45 32 100 54 98 77 12 36 54 98

# OUTPUT

Range       Frequency
=====       =========
0-40            3
41-60			3
61-80			1
81-100			3
*/
package main

import "fmt"

const students = 10 // constant representing the number of students

func main() {

	var marks = [students]int{} // array to store marks of students

	fmt.Println("Enter the marks obtained by students:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&marks[i]) // read marks from input
	}

	// variables to count frequency of marks in different ranges
	var r1 = 0 // 0-40
	var r2 = 0 // 41-60
	var r3 = 0 // 61-80
	var r4 = 0 // 81-100

	// iterate over marks array and count frequency in each range
	for i := 0; i < students; i++ {
		if marks[i] >= 0 && marks[i] <= 40 {
			r1++
		} else if marks[i] >= 41 && marks[i] <= 60 {
			r2++
		} else if marks[i] >= 61 && marks[i] <= 80 {
			r3++
		} else {
			r4++
		}
	}

	// print frequency of marks in different ranges
	fmt.Println("Range\t\tFrequency")
	fmt.Printf("0-40\t\t%d\n", r1)
	fmt.Printf("41-60\t\t%d\n", r2)
	fmt.Printf("61-80\t\t%d\n", r3)
	fmt.Printf("81-100\t\t%d\n", r4)
}
