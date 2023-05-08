package main

import (
	"fmt"
	"kodenesia-answer/service"
)

func main() {
	// Call Service (Solution) Here
	// Example answer := service.FunctionName(func params...)

	s := service.LampStatus(10, 4)
	fmt.Println(s)
}
