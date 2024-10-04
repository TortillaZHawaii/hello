package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}

	for _, num := range arr {
		if _, err := fmt.Println("num:", num); err != nil {
			err := errors.Join(errors.New("couldn't print"), err)
			panic(err)
		}
	}

	fmt.Println("hello from the other side")
}
