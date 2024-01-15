package main

import "fmt"

func PrintIo() { // 開頭要大寫才能被引用
	input := 0
	fmt.Print("請輸入數字:")
	fmt.Scanln(&input)
	fmt.Printf("你輸入的是:%d\n", input) // 格式化輸出
	//fmt.Println("你輸入的是:", input)
	fmt.Scanln()
}
