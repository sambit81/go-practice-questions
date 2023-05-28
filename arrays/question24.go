/*
Store name and mark of 10 students then display the top 5 names based on their performance.

Example:

# INPUT

Names:

Alice
Bob
Claire
David
Emily
Frank
Grace
Harry
Isabel
Jack

Marks:

85
92
78
91
89
95
82
84
90
84

Top 5 Students:
Frank - 95
Bob - 92
David - 91
Isabel - 90
Emily - 89
*/
package main

import (
	"fmt"
)

const size int = 10 // Constant representing the size of the array

func main() {
	// Declare marks array
	var marks [size]int

	// Declare names array
	var names [size]string

	fmt.Println("Enter names of students:")

	// Input names
	for i := 0; i < size; i++ {
		fmt.Scan(&names[i])
	}

	fmt.Println("Enter marks of students")

	// Input marks
	for i := 0; i < size; i++ {
		fmt.Scan(&marks[i])
	}

	// Sort the marks and names array according to marks scored
	for i := 0; i < size; i++ {
		// Compare each index with all the indices starting from the next index of current index
		for j := i + 1; j < size; j++ {
			if marks[i] < marks[j] {
				// If current index value is smaller than its next index then swap
				d := marks[i]
				marks[i] = marks[j]
				marks[j] = d

				// Names has to be swaped if marks are swaped
				s := names[i]
				names[i] = names[j]
				names[j] = s
			}
		}
	}

	// After sorting marks array will be in descending order and also the names array will be sorted accordingly
	// So, print the first five names
	fmt.Println("Top 5 Students:")
	for i := 0; i < 5; i++ {
		fmt.Println(names[i], "-", marks[i])
	}
}
