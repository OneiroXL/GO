package main

import (
	"fmt"
)

func main() {
	var chann chan int
	chann = make(chan int , 3)

	//管道的长度和容量
	fmt.Printf("管道的长度:%v,管道的容量:%v \n",len(chann),cap(chann))

	//加入数据
	chann <- 10
	chann <- 20

	//取出数据
	fmt.Println("取出的数据为:",<-chann)

	//管道的关闭(管道关闭只能读不能写)
	close(chann)

	//遍历管道(遍历管道时 要关闭管道)
	//不关闭管道遍历会出现deadlock错误
	for v := range chann {
		fmt.Println("值为:",v)
	}

	//声明只写管道
	var channWrite chan<- int = make(chan int,3)
	channWrite<-1
	
	//声明只读管道
	var channRead <-chan int = make(chan int,3)
	<-channRead

	for {
		//如果管道没有关闭也一直阻塞而deadlock
		select{
			case v := <-chann:{
				fmt.Println("取出的数据为:",v)
			}
			case v:= <-channRead:{
				fmt.Println("取出的数据为:",v)
			}
			default:{
				fmt.Println("两个都取不出数据")
				break
			}
		}
	}
}


func Write(chann chan<- int,v int){
	chann<-v
}

func Read(chann <-chan int) int{
	return <-chann
}