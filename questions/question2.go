/*
Store and display 10 numbers in one column and their square in another column.

Example:

Numbers         Square
1                1
2                4
3                9
4                16
5                25
6                36
7                49
8                64
9                81
10               100
*/
package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println("Enter 10 numbers")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Numbers\t\tSquare")
	for _, val := range arr {
		fmt.Println(val, "\t\t", val*val)
	}
}
