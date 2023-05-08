package main

import (
	"fmt"
	"kodenesia-answer/service"
)

func main() {
	// Call Service (Solution) Here
	// Example 		=> s := service.FunctionName(func params...)
	// Print Output => fmt.Println(s)

	s := service.FlyingRugs("100 90 30 2 85 80 200 33 54")
	fmt.Println(s)
}
