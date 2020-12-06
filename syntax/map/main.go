package main

import (
	"fmt"
	"sort"
)

func main() {
	// mapが参照型である事の確認
	m := make(map[string]string)
	// main関数内でmは変更していないが参照渡しされるので変更される。
	addMap(m)
	fmt.Println(m)
	// copyしているので、mに影響はない
	addMapWithCopy(m)
	fmt.Println(m)

	// ---------
	// mapに対するfor
	// 学籍番号と学生名のMap
	studentIDMap := map[int]string{
		3: "田中",
		1: "伊藤",
		2: "佐藤",
		4: "佐々木",
	}

	studentIDSlice := make([]int, len(studentIDMap))
	index := 0
	for key := range studentIDMap {
		studentIDSlice[index] = key
		index++
	}
	sort.Slice(studentIDSlice, func(i, j int) bool { return studentIDSlice[i] < studentIDSlice[j] })

	for _, key := range studentIDSlice {
		// fmt.Printfでフォーマットに従った文字列を標準出力に出せる
		fmt.Printf("Name of StudentID:%d is %s\n", key, studentIDMap[key])
	}
}

func addMap(m map[string]string) {
	m["a"] = "あ"
}
func addMapWithCopy(m map[string]string) {
	copied := make(map[string]string)
	for k, v := range m {
		copied[k] = v
	}
	copied["i"] = "い"
}
