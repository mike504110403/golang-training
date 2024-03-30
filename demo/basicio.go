package demo

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintIo() { // 開頭要大寫才能被引用
	input := 0
	fmt.Print("請輸入數字:")
	fmt.Scanln(&input)
	fmt.Printf("你輸入的是:%d\n", input) // 格式化輸出
	//fmt.Println("你輸入的是:", input)
	fmt.Scanln()
}

func PrintCalculateText() []int {
	// 任意長度陣列
	inputs := []int{}
	fmt.Print("請輸入數字（輸入空行結束）:\n")

	for {
		var inputStr string
		fmt.Scanln(&inputStr)
		// 輸入空值結束循環
		if inputStr == "" {
			break
		}

		// input轉int
		// TrimSpace: 字串去頭尾空白
		// strconv: string convert; Atoi(): ASCII to integer
		input, err := strconv.Atoi(strings.TrimSpace(inputStr))
		if err != nil {
			fmt.Println("輸入無效，請輸入數字！")
			continue
		}

		// 加入陣列
		inputs = append(inputs, input)
	}
	return inputs
}
