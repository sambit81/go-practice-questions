/*
Write a Go program that stores the names and marks of 5 students and displays the names in
descending order of their marks using the bubble sort algorithm.

Your program should prompt the user to enter the names of the 5 students and their respective marks,
and then display the sorted list of names with their corresponding marks.

Example:

# INPUT

Sambit: 54
Akansha: 98
Adarsh: 12
Samyam: 65
Somyam: 45

# OUTPUT

Names           Marks
Akansha         98
Samyam          65
Sambit          54
Somyam          45
Adarsh          12
*/
package main

import "fmt"

const num = 5 // Number of students

func main() {

	// Declare two arrays, one for names and another for marks
	var names = [num]string{}
	var marks = [num]int{}

	fmt.Printf("Enter the names of %d students:\n", num)

	// Read the names of students from input
	for i := 0; i < num; i++ {
		fmt.Scan(&names[i])
	}

	fmt.Println("Enter their marks:")

	// Read the marks of students from input
	for i := 0; i < num; i++ {
		fmt.Scan(&marks[i])
	}

	// Sort the names and marks array in descending order of marks
	for i := 0; i < num; i++ {
		for j := 0; j < num-i-1; j++ {
			if marks[j] < marks[j+1] {
				// Swap the marks
				m := marks[j]
				marks[j] = marks[j+1]
				marks[j+1] = m

				// Alternative to above swap
				// marks[j], marks[j+1] = marks[j+1], marks[j]

				// Swap the names
				n := names[j]
				names[j] = names[j+1]
				names[j+1] = n

				// Alternative to above swap
				// names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}

	// Display the sorted names and marks
	fmt.Println("Names\t\tMarks")
	for i := 0; i < num; i++ {
		fmt.Printf("%s\t\t%d\n", names[i], marks[i])
	}
}
