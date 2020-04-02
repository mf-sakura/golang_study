package main

import (
	"fmt"
	"os"
)

func main() {
	if err := CatFile("test.txt"); err != nil {
		fmt.Println(err)
	}
}

func CatFile(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("File open error: %v\n", err)
		return
	}

	defer func() {
		if err != nil {
			// エラー時にのみdeferで行いたい処理が書ける
			fmt.Println("Error Handling in defer called.")
		}
		// fileはCloseする必要がある。
		err = file.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println("File read error: ", err)
			return err
		}
		fmt.Print(string(buf[:n]))
	}

	// これだとエラー時にfileが開きっぱなしになる。
	// file.Close()

	return nil
}
