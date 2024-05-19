package demo

import (
	"fmt"
	"time"
)

func TimerDemo() {
	fmt.Println("now: ", time.Now().Format("2006-01-02 15:04:05"))

	timer := time.NewTimer(3 * time.Second)
	c := <-timer.C // 通道接收值
	fmt.Println(c)

	fmt.Println("now: ", time.Now().Format("2006-01-02 15:04:05"))
}

func TickerDemo() {
	done := make(chan struct{})
	go ticker(done)
	<-done
}

func ticker(done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan interface{})
	defer ticker.Stop()
	defer close(done)
	for {
		select {
		case c := <-ticker.C:
			fmt.Println("嚴心誼：歪你　現在時間: ", c.Format("2006-01-02 15:04:05"))
		case quit <- ticker.C:
			return
		}
	}
}
