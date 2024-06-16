package demo

import "fmt"

func VisitPrint(list []int, f func(int)) {
	for _, n := range list {
		f(n)
	}
}

func TestCallBack() {
	sli := []int{1, 5, 8, 6}
	VisitPrint(sli, func(value int) {
		fmt.Println(value)
	})
}
