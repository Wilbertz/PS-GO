package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("./go.mod")

	if err != nil {
		fmt.Println("This is the error code", err)
	}
	fmt.Println("Ok", err)
}
