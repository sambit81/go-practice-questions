/*
Store 11 to 19 in reverse order in an array and then display those numbers in ascending order.

Example:

11
12
13
14
15
16
17
18
19
*/
package main

import "fmt"

func main() {
	var arr [9]int
	for i := 0; i < 9; i++ {
		arr[i] = 19 - i
	}
	fmt.Println("Values in reverse order:", arr)
	fmt.Println("Values in correct order:")
	for i := 8; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}
