package main

import "fmt"

func main() {
	dddLengthMins := 275
	cawLengthMins := 30

	if dddLengthMins > cawLengthMins {
		fmt.Println("Docker Deep Dive is longer than Containers on AWS Wavelength")
	}
}
