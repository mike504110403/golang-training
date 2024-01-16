package main // 可執行程式封包要用main

import (
	"fmt" // 載入內建封包 輸出輸入用
	"reflect"
)

// 撰寫程式 > 建置 > 執行
// 建置: go build 檔案名稱 > 生出執行檔 windows是exe
// 執行: 檔名 + .\代表信任檔案

// main函式 程式進入點
func main() {
	//fmt.Println("Hello Golang")
	//varDeclare() // 呼叫func
	//setVar()            // 設定變數
	//fmt.Scanln(&intVar) // 放到變數中 &變數名稱: 取得變數的指標
	//printOutVar(intVar) // print變數 (可依照傳入參數做額外操作)
	//PrintIo()
	fmt.Printf("計算結果為:%d", (MutiAdd(PrintCalculateText())))

}

// 變數 這邊是全域
// note 型別如同C# 為強型別
var stringVar string
var 中文變數 string    // 中文也可宣告
var intVar int = 4 // 宣告時可以同時設定值
var float float64 = 3.1415
var boolVar bool = false
var charVar rune = 'a' // 字符 背後用整數紀錄

func setVar() {
	stringVar = "字串"
	中文變數 = "中文變數"
}

func printOutVar(input interface{}) {
	// 透過型別斷言記述入的型別
	switch inputype := input.(type) {
	case string:
		fmt.Println(reflect.TypeOf(inputype), input) // reflect.TypeOf() 是型別
	case int:
		fmt.Println(reflect.TypeOf(input), input) //
		// 可再放其他操作
	}
	// note.  inputype 實際上記到的值 同input 但可以拿來斷言
	// 這樣不是常量 不能放case
	/*valueType := reflect.TypeOf(input) := input.(type) 型別斷言不是
	switch valueType{
	case string:
		fmt.Println(valueType, input)
	}*/
}
