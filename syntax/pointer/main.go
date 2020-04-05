package main

import (
	"fmt"
)

func main() {
	n := 100
	// 値渡し
	// コピーされるので、元のnに変化はない
	returnValueA := increment(n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Return Value of increment is %d\n", returnValueA)
	fmt.Printf("Same n and returnValueA addresses? %v\n", compareAddress(&n, &returnValueA))
	// 参照渡し
	// 戻り値を受け取らなくても渡した変数が書き換わる
	// & は変数のアドレスを取得する
	returnValueB := incrementWithPointer(&n)
	fmt.Printf("Value of n is %d\n", n)
	fmt.Printf("Same n and returnValueB addresses? %v\n", compareAddress(&n, &returnValueB))
}

func increment(n int) int {
	return n + 1
}

// アドレスから変数の値にアクセスするには * をつける
func incrementWithPointer(n *int) int {
	*n++
	// 戻り値はint型なので変数の値にアクセスする
	return *n
}

func compareAddress(a, b *int) bool {
	if *a == *b {
		return true
	} else {
		return false
	}
}
