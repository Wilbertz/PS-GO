package main

import "fmt"

func main() {

	coursesInProg := []string{
		"Docker & Kubernetes",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive"}

	for i, data := range coursesInProg {
		fmt.Println(i, data)
	}
}
