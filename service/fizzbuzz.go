package service

import "fmt"

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i)

		if i%3 == 0 && i%5 == 0 {
			fmt.Print(" FizzBuzz")
		} else if i%3 == 0 {
			fmt.Print(" Fizz")
		} else if i%5 == 0 {
			fmt.Print(" Buzz")
		}
		fmt.Println()
	}
}
