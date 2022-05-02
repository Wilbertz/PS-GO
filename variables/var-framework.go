package main

import "fmt"

func main() {
	name := "Nigel Poulton"
	course := "Getting started with Kubernetes"

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)

	fmt.Println("Your are currently watching", course)
}

func updateCourse(course *string) string {
	*course = "Getting Started with Docker"
	fmt.Println("Update course to", *course)
	return *course
}
