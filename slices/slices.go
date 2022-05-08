package main

import "fmt"

func main() {
	courses := make([]string, 5, 10)

	courses[0] = "Docker & Kubernetes"
	courses[1] = "Getting started with Docker"
	courses[2] = "Getting started with Kubernetes"

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))
}
