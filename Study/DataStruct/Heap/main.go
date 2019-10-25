package main

import (
	"./maxHeap"
	"./minHeap"
	"time"
	"math/rand"
	"fmt"
)

func main()  {
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包

	maxHeap := maxHeap.NewMaxHeap(10)
	minHeap := minHeap.NewMinHeap(10)

	for index := 0; index < 10; index++ {
		var num int = rand.Intn(100)
		maxHeap.Push(num)
		minHeap.Push(num)
	}


	for index := 0; index < 10; index++ {
		v,_ := maxHeap.Pop();
		fmt.Println(v) 
	}
	fmt.Println("-----------------") 
	for index := 0; index < 10; index++ {
		v,_ := minHeap.Pop();
		fmt.Println(v) 
	}

}