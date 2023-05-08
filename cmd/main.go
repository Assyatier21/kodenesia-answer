package main

import (
	"fmt"
	"kodenesia-answer/service"
)

func main() {
	// Call Service (Solution) Here
	// Example answer := service.FunctionName(func params...)

	s := service.FingersCount(1, 10)
	fmt.Println(s)
}
