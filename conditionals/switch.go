package main

import "fmt"

func main() {
	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case1")
	case "Kubernetes":
		fmt.Println("Case2")
	case "K8s":
		fmt.Println("Case3")
	case "Istio":
		fmt.Println("Case4")
	default:
		fmt.Println("Case5")
	}
}
