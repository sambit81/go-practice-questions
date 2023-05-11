/*
Store 10 numbers in an array. Then separate only palindromes from the array and display those numbers.

Example:

# INPUT

111
21
26
616
888
322
122
222
000
120

# OUTPUT

111 616 888 222 0
*/
package main

import (
	"fmt"
	"strconv" // import the "strconv" package to convert integers to strings
)

func main() {

	var array [10]int // declare an array of size 10 to store the input numbers

	fmt.Println("Enter 10 numbers:")

	// loop to take input of 10 numbers and store them in the array
	for i := 0; i < 10; i++ {
		fmt.Scan(&array[i])
	}

	var palindromes []int // declare a slice to store the palindrome numbers

	// loop to check each number in the array for palindrome
	for _, num := range array {
		var strformat = strconv.Itoa(num) // convert the integer to a string using "strconv.Itoa"
		var isPalin = true                // assume the number is a palindrome

		// loop to check if the number is a palindrome
		for i := 0; i < len(strformat)/2; i++ {
			if strformat[i] != strformat[len(strformat)-i-1] { // check if the first and last characters of the string match, then second and second-to-last characters, and so on.
				isPalin = false // if any pair of characters don't match, the number is not a palindrome
				break           // exit the loop since we already know the number is not a palindrome
			}
		}

		if isPalin == true { // if the number is a palindrome
			palindromes = append(palindromes, num) // append it to the "palindromes" slice
		}
	}

	fmt.Println("Palindromes found:", palindromes) // display the "palindromes" slice containing all palindrome numbers
}
