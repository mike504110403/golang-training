package main

import (
	"strconv"
	"strings"
)

// 兩參數相加
func Add(par1, par2 int) int {
	return typeAssert(par1) + typeAssert(par2)
}

// 多參數相加
func MutiAdd(pars []int) int {
	result := 0
	// 忽略索引瀝遍
	for _, par := range pars {
		result += typeAssert(par)
	}

	return result
}

// 型別斷言 todo: 轉float 處理小數點
func typeAssert(par interface{}) int {
	switch v := par.(type) {
	case int:
		return v
	case string: // 字串轉int後回傳 失敗回傳0
		convert, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return 0
		}
		return convert
	default:
		return 0
	}
}
