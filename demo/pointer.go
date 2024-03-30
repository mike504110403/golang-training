package demo

import "fmt"

func Pointer() {
	str := "字串的值" // 一般變數
	str2 := "換位置"
	var ptr *string // 指標變數
	var strptr string
	ptr = &str // 取址給到指標變數
	strptr = *ptr
	//strptr = &str // 報錯: 非指標變數不能記位址

	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(strptr)
	fmt.Println(&strptr)
	ptr = &strptr
	fmt.Println(ptr)
	fmt.Println("---------------")
	ptr = &str2
	fmt.Println(ptr)
}

// 透過指標在記憶體位址不動的情況交換值
func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
