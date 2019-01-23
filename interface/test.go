package main

import "fmt"

type Double interface {
	func1(int) int
	func2(int) int
}

type impDouble struct {
	a int
}

func (i impDouble) func1(a int) int {
	fmt.Println("func1", a)
	return a
}

func (i *impDouble) func2(a int) int {
	fmt.Println("func2", a)
	b := a
	return b
}

func (i *impDouble) String() string {
	return fmt.Sprint("impDouble", i.a)

}

func main() {

	var a = impDouble{a: 1}
	var double Double = &a
	double.func1(1)
	double.func2(2)
	fmt.Println(double)
	fmt.Println(a)

}
