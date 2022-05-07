package main

import "fmt"

func main() {

	coursesInProg := []string{
		"Docker & Kubernetes",
		"Docker Networking",
		"Getting Started with Kubernetes",
		"Kubernetes Deep Dive"}

	coursesCompleted := []string{
		"Docker & Kubernetes",
		"Docker Deep Dive"}

	for i, data := range coursesInProg {
		fmt.Println(i, data)
		for _, j := range coursesCompleted {
			if data == j {
				fmt.Println("Hey we found a clash.", j, "is in both lists")
			}
		}
	}
}
