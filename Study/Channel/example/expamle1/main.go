
package main

import (
	"fmt"
)

type Result struct{
	Key int
	Value int
}


var (
	chanNum chan int = make(chan int,16)
	chanRes chan Result = make(chan Result,500)
	chanExit chan<- bool = make(chan bool,8)
)

func Write(){
	for i := 1; i <= 500; i++ {
		chanNum<-i	
	}
	close(chanNum)
}

func AddUp(){
	for i := 0; i < 500; i++ {
		num,ok := <-chanNum
		if(!ok){
			chanExit<-true
			break;
		}
		var res int
		for i := 1; i <= num; i++ {
			res = res + i
		}
		//放入结果管道
		var r Result = Result{}
		r.Key = num;
		r.Value = res;
		chanRes <- r;
	}
}

func ShowRes(){
	for r := range chanRes {
		fmt.Printf("结果为:res[%v] = %v \n",r.Key,r.Value)
	}
}


func main()  {
	go Write()
	for i := 0; i < cap(chanExit); i++ {
		go AddUp()
	}

	for {
		if(len(chanExit)== cap(chanExit)){
			close(chanRes)
			close(chanExit)
			break
		}
	}
	ShowRes()

	fmt.Println("执行完毕!!!")
}