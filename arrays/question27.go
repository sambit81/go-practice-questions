/*
A phonebook is already created. Ask one name as input then display the telephone number of that person if found. Otherwise display proper message.

Example:

# INPUT

Enter a name to search: Jane Smith

# OUTPUT

Telephone number of Jane Smith is 9876543210
===============================================================
# INPUT

Enter a name to search: Jane

# OUTPUT

No telephone number found for Jane
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const size int = 5 // Constant representing the size of the array

func main() {
	// Create a map to store name and telephone number pairs
	phoneBook := make(map[string]string)

	// Store name and telephone number of 5 people
	phoneBook["John Doe"] = "1234567890"
	phoneBook["Jane Smith"] = "9876543210"
	phoneBook["Mike Johnson"] = "5555555555"
	phoneBook["Emily Brown"] = "1111111111"
	phoneBook["Alex Davis"] = "9999999999"

	fmt.Print("Enter a name to search: ")

	// Create a reader to read input from standard input
	reader := bufio.NewReader(os.Stdin)

	// Read the input name and remove leading/trailing whitespace
	search_name, _ := reader.ReadString('\n')
	search_name = strings.TrimSpace(search_name)

	// Lookup telephone number based on the input name
	telephoneNumber, found := phoneBook[search_name]
	if found {
		fmt.Printf("Telephone number of %s is %s\n", search_name, telephoneNumber)
	} else {
		fmt.Printf("No telephone number found for %s\n", search_name)
	}
}
