package main

import {
	"fmt"
	"sync"
	"time"
}
// thread控制
var wg sync.WaitGroup
// 鎖變數控制 race condition
var mu sync.Mutex

func foo() {
    for i := 1; i <= 10; i++ {
        fmt.Println("Foo: ", i)
    }
	// 結束該執行續
    wg.Done()
}

func bar() {
    for i := 1; i <= 10; i++ {
        fmt.Println("Foo: ", i)
    }
    wg.Done()
}

func main() {
	// 執行續容量
    wg.Add(2)
    go foo()
    go bar()
	// 等背景執行續跑完
    wg.Wait()
}

func withdraw() {
    mu.Lock()
    balance := money
    time.Sleep(3000 * time.Millisecond)
    balance -= 1000
    money = balance
    mu.Unlock()
    fmt.Println("After withdrawing $1000, balace: ", money)
    wg.Done()
}

func main() {
    fmt.Println("We have $1500")
    wg.Add(2)
    go withdraw() // first withdraw
    go withdraw() // second withdraw
    wg.Wait()
}