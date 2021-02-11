package main

import "fmt"

func ShowFizzBuzz(numbers []int) {

	for _, val := range numbers {
		if val%3 == 0 {
			fmt.Println("Fizz")
		} else if val%5 == 0 {
			fmt.Println("Buzz")
		} else if val%15 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(val)
		}
	}
}

func main() {
	numbers := make([]int, 100)

	for i, _ := range numbers {
		numbers[i] = i + 1
	}

	ShowFizzBuzz(numbers)
}
