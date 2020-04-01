package main

import (
	"fmt"
)

func main() {
	n := 100
	// 値渡し
	// コピーされるので、元のnに変化はない
	returnValue := increment(n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Return Value of increment is %d\n", returnValue)
	fmt.Printf("[main]Pointer Value of n is %d\n", &n)

	// 参照渡し
	// 戻り値を受け取らなくても渡した変数が書き換わる
	incrementWithPointer(&n)
	fmt.Printf("Value of n is %d\n", n)
}

func increment(n int) int {
	fmt.Printf("[increment]Pointer Value of n is %d\n", &n)
	return n + 1
}
func incrementWithPointer(n *int) {
	fmt.Printf("[incrementWithPointer]Pointer Value of n is %d\n", n)
	*n++
}
