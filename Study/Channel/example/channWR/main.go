package main

import (
	"fmt"
)


var (
	channInt chan int = make(chan int , 50)
	channIsOver chan bool = make(chan bool,1)
)

func main()  {
	go Write()
	go Read()

	for{
		if(<-channIsOver){
			return
		}
	}
}


func Write(){
	for i := 0; i < 50; i++ {
		channInt <- i
	}
	close(channInt)
}

func Read(){
	for i := 0; i < 50; i++ {
		fmt.Println(<-channInt)
	}
	channIsOver <- true
	close(channIsOver)
}