package main

import "fmt"

func main() {
	// lengthは3
	a := make([]int, 3, 3)
	fmt.Printf("Length of a is %d\n", len(a))
	// lengthが0
	b := make([]int, 0, 3)
	fmt.Printf("Length of b is %d\n", len(b))

	// Sliceはポインタを含むので、`=`で代入してしまうと同じアドレスを使い回す事になる
	// これは関数に渡す場合でも同じ
	c := a
	a[0] = 2
	fmt.Printf("c[0] is %d\n", c[0])
	fmt.Printf("Address of a[0] is %v\n", &a[0])
	fmt.Printf("Address of c[0] is %v\n", &c[0])

	// Sliceのコピー
	d := make([]int, 3, 3)
	// 要素数が異なる場合は、2Sliceで最小の要素数だけコピーされる。
	copy(d, a)
	// 1要素目がコピーされている事を確認する
	fmt.Printf("d[0] is %d\n", d[0])
	// ポインタを共有していない事を確認する
	a[1] = 3
	fmt.Printf("c[0] is %d\n", d[1])
	fmt.Printf("Address of a[1] is %v\n", &a[1])
	fmt.Printf("Address of d[1] is %v\n", &d[1])

	// index out of rangeになるとpanicする
	fmt.Println(a[4])
}
