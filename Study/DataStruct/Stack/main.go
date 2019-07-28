package main


import (
	"fmt"
	"./stack"
)


func main() {
	stack := stack.NewStack(10)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push("6 - 6")
	stack.Push(7)
	stack.Push("8 - 8")
	stack.Push(9)
	stack.Push(10)

	fmt.Println(stack.Push(11))

	stack.Show()

	v,_ := stack.Pop()
	fmt.Println(v)

	stack.Show()
}