package main

import (
	"./queue"
	"fmt"
)

func main()  {

	queue := queue.NewQueue(5)

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)

	queue.Show()

	fmt.Println(queue.Push(6))

	v,_ := queue.Pop()
	fmt.Println(v)

	v4,_ :=  queue.Pop()
	fmt.Println(v4)

	queue.Push(7)

	v3,_ :=  queue.Pop()
	fmt.Println(v3)

	queue.Show()

	queue.Pop()
	queue.Pop()
	queue.Pop()

	_,err := queue.Pop()
	fmt.Println(err)

	queue.Push("我是我")

	v5,_:= queue.Pop()
	fmt.Println(v5)
}