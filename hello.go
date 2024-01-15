package main // 可執行程式封包要用main
import "fmt" // 載入內建封包 輸出輸入用

// main函式 程式進入點
func main(){
	fmt.Println("Hello Golang")
	stringDeclare()
}
// 撰寫程式 > 建置 > 執行
// 建置: go build 檔案名稱 > 生出執行檔 windows是exe
// 執行: 檔名 + .\代表信任檔案

/*變數宣告*/
// 全域宣告
var stringVar string
var 中文變數 string // 中文也可宣告
var boolVar bool = false // 宣告時可以同時設定值
var charVar rune = 'a' // 字符 背後用整數紀錄
func stringDeclare(){
	stringVar = "字串"
	中文變數 = "中文變數"

	fmt.Println(stringVar)
	fmt.Println(中文變數)
	fmt.Println(boolVar)
	fmt.Println(charVar)
}