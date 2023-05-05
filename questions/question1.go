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
