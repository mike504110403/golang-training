package demo

type Num int

func (n Num) Equal(i Num) bool {
	return n == i
}

func (n Num) LessThan(i Num) bool {
	return n < i
}

func (n Num) GreaterThan(i Num) bool {
	return n > i
}

func (n *Num) Mutiple(i Num) {
	*n = *n * i
}

func (n *Num) Divide(i Num) {
	*n = *n / i
}

type NumI interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	GreaterThan(i Num) bool
	Mutiple(i Num)
	Divide(i Num)
}

var x Num = 8
var y NumI = &x
