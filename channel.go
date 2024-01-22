package main

import (
	"fmt"
)

// channel
func channelRecieve() {
	// 建立通道變數
	c := make(chan int)

	// 在goroutine之下執行這個function
	go func() {
		for i := 0; i < 10; i++ {
			// 將值指給通道
			c <- i
		}
		// 關通道 關了後只能讀取 不能接收
		close(c)
	}()
	// 逐筆秀c中的值
	for n := range c {
		fmt.Println(n)
	}
}

func channelSend() {
	// 發送用
	c := make(chan int)
	// 接受用
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		defer close(c) // 延遲執行，確保在函數退出時關閉 c

		doneCounter := 0
		for range done { // done接收兩次時關閉 c
			doneCounter++
			if doneCounter == 2 {
				return
			}
		}
	}()

	for n := range c {
		fmt.Println(n)
	}
}
