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

	for _, k := range sortByKeys(studentIDMap) {
		// fmt.Printfでフォーマットに従った文字列を標準出力に出せる
		fmt.Printf("Name of StudentID:%d is %s\n", k, studentIDMap[k])
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

func sortByKeys(m map[int]string) (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}
	// 期待されたソートの仕方ではないがsort.Sliceの第一引数に
	// mapだったものをどうsliceにするか？が解決できなかった
	sort.Ints(keys)
	//fmt.Println(keys)
	return
}
