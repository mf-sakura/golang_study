package main

// Sum is a function for adding x to y
func Sum(x int, y int) int {
	sum := x + y
	// 不自然だけどcoverageを見たいので条件分岐させる
	if sum >= 1000 {
		return 0
	}
	return sum
}

func main() {
	Sum(5, 5)
}
