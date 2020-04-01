package main

import "fmt"

func main() {

	numbers := make([]*int, 0, 3)
	for i := 0; i < 3; i++ {
		key := i
		numbers = append(numbers, &key)
	}

	fmt.Println("Values:", *numbers[0], *numbers[1], *numbers[2])
	fmt.Println("Addresses:", numbers[0], numbers[1], numbers[2])

}
