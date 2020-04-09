package main

import (
	"fmt"
	"strconv"
)

// Sum is a function for adding x to y
func Sum(x int, y int) int {
	sum := x + y
	// 不自然だけどcoverageを見たいので条件分岐させる
	if sum >= 1000 {
		return 0
	}
	return sum
}

// Counter is a function for counting x
func Counter(x int) string {
	if x > 99 {
		return "99+"
	}
	return strconv.Itoa(x)
}

func main() {
	sum := Sum(5, 5)
	fmt.Printf("%s\n", Counter(sum))
}
