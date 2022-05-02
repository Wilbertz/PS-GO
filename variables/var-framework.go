package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
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
