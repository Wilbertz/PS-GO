package main

import "fmt"

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float64
	}
	gettingStartedWithK8s := courseMeta{
		author: "Nigel Poulton",
		level:  "Intermediate",
		rating: 5,
	}
	fmt.Println("Author of Getting started with Kubernetes is:", gettingStartedWithK8s.author)
}
