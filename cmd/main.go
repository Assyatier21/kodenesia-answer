package main

import (
	"fmt"
	"kodenesia-answer/service"
)

func main() {
	s := service.FindBrackets("{[()]}")
	fmt.Println(s)
}
