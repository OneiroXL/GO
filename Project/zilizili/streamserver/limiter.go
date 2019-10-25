package streamserver

import (
	"log"
)

type Limiter struct {
	currentConCount int
	limitChan chan int
}


/*
构造函数
*/
func NewLimiter(n int) *Limiter {
	limit := &Limiter{
		currentConCount:n,
		limitChan : make(chan int, n),
	}
	for i := 0; i < n; i++ {
		limit.limitChan <- 1
	}
	return limit
}

/*
归还
*/
func (this *Limiter) Return() {
	log.Println("Return")
	this.limitChan <- 1
}

/*
借出
*/
func (this *Limiter) Borrow() bool {
	log.Println("Borrow")
	if(len(this.limitChan) == 0){
		log.Printf("Reached the rate limition")
		return false
	}
	<-this.limitChan  
	return true
}