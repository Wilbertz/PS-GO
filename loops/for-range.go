package main

import "fmt"

func main() {

	coursesInProg := []string{
		"Docker & Kubernetes",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive"}

	for _, i := range coursesInProg {
		fmt.Println(i)
	}
}
