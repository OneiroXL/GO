package main

import (
	"fmt"
)

func Div(num1,num2 int) int {
	//使用defer recover 来捕获异常
	defer func(){
		err := recover() //recover() 可以捕获异常
		if(err != nil){
			fmt.Println("错误:",err)
		}
	}()
	res := num1 / num2
	return res
}


func main(){
	res := Div(10,0)
	fmt.Println(res)
}