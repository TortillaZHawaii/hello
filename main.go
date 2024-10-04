package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	for _, num := range numbers {
		if err := printNumber(num); err != nil {
			panic(fmt.Errorf("failed to print number: %w", err))
		}
	}

	fmt.Println("hello from the other side")
}

func printNumber(num int) error {
	_, err := fmt.Println("num:", num)
	if err != nil {
		return fmt.Errorf("couldn't print: %w", err)
	}
	return nil
}
