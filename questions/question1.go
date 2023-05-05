/*
*
To display 10 numbers stored in an array and their index position in two adjacent
columns with appropriate headings.

Example:

Numbers         Index Position
2                0
5                1
9                2
5                3
10               4
9                5
87               6
56               7
12               8
63               9
*/
package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println("Enter 10 numbers:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Numbers\t\tIndex Position")
	for index, val := range arr {
		fmt.Println(val, "\t\t", index)
	}
}
