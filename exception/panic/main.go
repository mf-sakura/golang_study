package main

import "fmt"

func main() {
	// Sliceのlengthを超えてアクセスするとpanic
	indexOutOfRangeExample()

	// nil pointerに対してfieldやメソッドを得ようとしようとするpanic
	nilPointerAccess()

	// 初期化していないMapに書き込みをするとpanic
	uninitializeMap()

	// panicでプロセスが途中終了する
	panic("panic")
	fmt.Println("Successfully Shutdown")
}

func indexOutOfRangeExample() {
	defer func() {
		// recover処理(execptionのcatchに相当)
		if r := recover(); r != nil {
			fmt.Printf("Recovered.\n %v\n\n", r)
		}
	}()
	slice := make([]int, 0, 5)

	for i := 1; i <= 6; i++ {
		slice[i] = i
	}
}

type myStruct struct {
	field string
}

func nilPointerAccess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered.\n %v\n\n", r)
		}
	}()

	var myPointer *myStruct
	fmt.Println(myPointer.field)
}

func uninitializeMap() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered.\n %v\n\n", r)
		}
	}()
	var m map[string]string

	m["key"] = "value"
}
