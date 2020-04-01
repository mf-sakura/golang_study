package main

import (
	"fmt"
	"os"
)

func main() {
	if err := catFile("test.txt"); err != nil {
		fmt.Println(err)
	}
}

func catFile(path string) (err error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		fmt.Printf("File open error: %v\n", openErr)
		err = openErr
		return
	}

	defer func() {
		if err != nil {
			// エラー時にのみdeferで行いたい処理が書ける
			fmt.Println("Error Handling in defer called.")
		}
		// fileはCloseする必要がある。
		// 本当はエラーハンドリングが必要(課題)
		if err := file.Close(); err != nil {
			fmt.Println("Error has occurred by file.Close().")
		}
	}()

	// //エラーを明示的に返してdeferが呼ばれるか確認する。
	// return errors.New("error")

	buf := make([]byte, 1024)
	for {
		n, readErr := file.Read(buf)
		if n == 0 {
			break
		}
		if readErr != nil {
			fmt.Println("File read error: ", err)
			err = readErr
			return
		}
		fmt.Print(string(buf[:n]))
	}

	// これだとエラー時にfileが開きっぱなしになる。
	// file.Close()

	return nil
}
