package main

import (
	"fmt"
	"sort"
)

func main() {
	// mapが参照型である事の確認
	// map == hash, map[KeyType]ValueType
	// nilを入れたmapを生成することはできるが、書き込みができない(panicが起こる)ので、初期化したい時はmakeを使う
	// makeで初期化できるのはmap, slice, channel
	m := make(map[string]string)
	// main関数内でmは変更していないが参照渡しされるので変更される。
	addMap(m)
	fmt.Printf("After addMap is %v\n", m)
	// copyしているので、mに影響はない
	addMapWithCopy(m)
	fmt.Printf("After addMapWithCopy is %v\n", m)

	// ---------
	// mapに対するfor
	// 学籍番号と学生名のMap
	studentIdMap := map[int]string{
		3: "田中",
		1: "伊藤",
		2: "佐藤",
		4: "佐々木",
	}

	// sliceを定義して、ここにkeyを入れていく
	studentIds := []int{}
	for key := range studentIdMap {
		studentIds = append(studentIds, key)
	}
	sort.Ints(studentIds)

	for _, k := range studentIds {
		fmt.Printf("Name of StudentID:%d is %s\n", k, studentIdMap[k])
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
