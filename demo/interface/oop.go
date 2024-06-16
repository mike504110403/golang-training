package oop

import (
	"golang-training/demo/interface/oop1"
	"golang-training/demo/interface/oop2"
)

type Num int

func (n Num) Equal(i int) bool {
	return int(n) == i
}

func (n Num) LessThan(i int) bool {
	return int(n) < i
}

func (n Num) GreaterThan(i int) bool {
	return int(n) > i
}

func (n Num) Sum(i int) int {
	return int(n) + i
}

// 因為NumInteface1和NumInterface2的方法相同
// Num同時時做了這兩個介面的實現
// 這樣就可以達到多型的效果

var f1 Num = 6
var f2 oop1.NumInterface1 = f1

// var f3 oop2.NumInterface2 = f2
// 1比2少了Sum方法，所以無法轉換1到2
// 但1為2的子集，所以可以轉換2到1
// 在這邊 Num 同樣做了兩個介面的實現
var f4 oop2.NumInterface2 = f1
var f5 oop1.NumInterface1 = f4
