/*
Store 15 Fibonacci numbers in the array and then display the numbers both in ascending as well as in descending order.

Fibonnaci numbers are numbers present in a sequence, where a particular number is formed by sum of the previous two numbers in the sequence.
The first two fibonacci numbers are fixed in the sequence which is 0 and 1. The third number will be 1 which is formed by sum of previous two numbers (0 and 1).
The fourth number will be 2 which is the sum of previous two fibonacci numbers (sum of 2nd and 3rd fibonacci numbers).

Example:

# INPUT

0 1 1 2 3 5 8 13 21 34 55 89 144 233 377

# OUTPUT

Ascending order: 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377
Descending order: 377 233 144 89 55 34 21 13 8 5 3 2 1 1 0
*/
package main

import "fmt"

const num = 15 // Define a constant variable for the number of Fibonacci numbers to generate

func main() {

	var arr = [num]int{} // Declare an integer array to store Fibonacci numbers

	arr[0] = 0 // First Fibonacci number is 0
	arr[1] = 1 // Second Fibonacci number is 1

	// Generate the remaining Fibonacci numbers
	for i := 2; i < num; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	// Print the Fibonacci numbers in ascending order
	fmt.Println("Fibonacci numbers in ascending order:")
	for i := 0; i < num; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()

	// Print the Fibonacci numbers in descending order
	fmt.Println("Fibonacci numbers in descending order:")
	for i := num - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
