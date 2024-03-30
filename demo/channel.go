package demo

import (
	"fmt"
	"sync"
)

// channel
func ChannelRecieve() {
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
		// 通道關閉後 會lock
		//c <- 20 // => panic: send on closed channel
	}()

	// 逐筆秀c中的值
	for n := range c {
		fmt.Println(n)
	}
}

func ChannelSend() {
	// 發送用
	c := make(chan int)
	var wg sync.WaitGroup

	// 啟動第一個thread
	wg.Add(1) // 給容量
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// 啟動第二個thread
	wg.Add(1) // 給容量
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// 啟動第三個thread
	wg.Add(1) // 給容量
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
	}()

	// 監控用
	//done := make(chan bool)
	go func() {
		defer close(c) // 延遲執行，確保在函數退出時關閉 c
		wg.Wait()      // 等待兩個thread完成
	}()

	wg.Wait()
}
