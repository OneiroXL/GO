package main

import(
	"fmt"
)

/*
闭包
*/

func AddUpper() func (int) int{
	var a int = 10
	return func (x int) int {
		a = a + x
		return a
	}
}

func main()  {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}