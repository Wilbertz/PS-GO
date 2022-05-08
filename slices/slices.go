package main

import "fmt"

func main() {
	courses := []string{"Docker & Kubernetes",
		"Getting started with Docker",
		"Getting started with Kubernetes"}

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))

	fmt.Println(courses)
}
