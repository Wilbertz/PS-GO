package main

import "fmt"

func main() {
	courses := make([]string, 5, 10)

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))
}
