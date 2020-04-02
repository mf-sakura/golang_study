package main

import "fmt"

func main() {
	// lengthは3
	a := make([]int, 3, 3)
	fmt.Printf("Length of a is %d\n", len(a))

	// Sliceはポインタを含むので、`=`で代入してしまうと同じアドレスを使い回す事になる
	// これは関数に渡す場合でも同じ
	b := a
	a[0] = 2
	fmt.Printf("b[0] is %d\n", b[0])
	fmt.Printf("Address of a[0] is %v\n", &a[0])
	fmt.Printf("Address of b[0] is %v\n", &b[0])

	// Sliceのコピー
	c := make([]int, 3, 3)
	// 要素数が異なる場合は、2Sliceで最小の要素数だけコピーされる。
	copy(c, a)
	// 1要素目がコピーされている事を確認する
	fmt.Printf("c[0] is %d\n", c[0])
	// ポインタを共有していない事を確認する
	a[1] = 3
	fmt.Printf("c[1] is %d\n", c[1])
	fmt.Printf("Address of a[1] is %v\n", &a[1])
	fmt.Printf("Address of c[1] is %v\n", &c[1])

	// index out of rangeになるとpanicする
	fmt.Println(a[4])
}
