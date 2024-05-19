package demo

import (
	"fmt"
)

func LoopDemo() {
	i := 0
	for {
		fmt.Printf("[outerloop]: %d\n", i)
		i++
		j := 0
		for {
			fmt.Printf("innerloop: %d\n", j)
			j++
			if j > 5 {
				break
			}
		}
		if i > 10 {
			break
		}
	}
}

func JumpLoopDemo() {
	i := 0
JumpLoop:
	for {
		//if i > 10 {
		//	break
		//}
		fmt.Printf("[outerloop]: %d\n", i)
		i++
		j := 0
		for {
			fmt.Printf("innerloop: %d\n", j)
			j++
			if j > 5 {
				if i > 10 {
					break JumpLoop // -> 終止該標籤對應的迴圈
				}
				break
			}
		}
	}
}

func Looprange() {
	str := "zdfgjhoajerg"
	strBytes := []byte(str)
	for i, char := range str {
		fmt.Printf("第 %d 個字元: 0x%x -> %c\n", i, char, char)
		char = 't'        // -> 這邊的char會是複本 改不到原本的str
		strBytes[i] = 'y' // -> 透過字串轉byte後 改指定位置的字元
	}
	fmt.Println(str)
	fmt.Println(string(strBytes))
}
