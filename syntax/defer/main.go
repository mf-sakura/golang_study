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

// CatFile file(test.txt)の中身を `cat` コマンドのように標準出力します
func catFile(path string) (err error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		fmt.Printf("File open error: %v\n", openErr)
		err = openErr
		return
	}

	// 課題検証用 defer実行前にClose()して必ず already closed を発生させる
	// file.Close()
	defer func() (err error) {
		if err != nil {
			// エラー時にのみdeferで行いたい処理が書ける
			fmt.Println("Error Handling in defer called.")
		}
		// fileはCloseする必要がある。
		// 本当はエラーハンドリングが必要(課題)
		if err = file.Close(); err != nil {
			// エラー時にのみdeferで行いたい処理が書ける
			fmt.Printf("Failed file.Close(). %v", err)
		}
		return
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
