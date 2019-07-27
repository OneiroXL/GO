package stack

import (
	"errors"
	"fmt"
)

var (
	max     int //最大容量
	top     int //栈顶
	dataArr []interface{}
)

type stack struct {
}

func NewStack(cap int) *stack {
	max = cap
	top = 0
	dataArr = make([]interface{}, cap, cap)
	return &stack{}
}

/*
进栈
*/
func (this *stack) Push(v interface{}) error {
	if this.IsFull() {
		return errors.New("Stack Is Full")
	}
	dataArr[top] = v
	top++
	return nil
}

/*
出栈
*/
func (this *stack) Pop() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Stack Is Empty")
	}
	top--
	return dataArr[top], nil
}

/*
是否满栈
*/
func (this *stack) IsFull() bool {
	return top == max
}

/*
是否空栈
*/
func (this *stack) IsEmpty() bool {
	return top == 0
}

/*
打印栈
*/
func (this *stack) Show() {
	for index := top - 1; index >= 0; index-- {
		fmt.Println(dataArr[index])
	}
}
