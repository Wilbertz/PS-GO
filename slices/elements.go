package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice[1] = 0
	fmt.Println(mySlice)
}
