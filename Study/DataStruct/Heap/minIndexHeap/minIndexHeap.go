package minHeap

import (
	"errors"
	"fmt"
)

type MinIndexHeap struct {
	heap [100]int
	index [100]int
	count int
	cap int
}

/*
创建一个最小堆
*/
func MinIndexHeap(cap int) *MinHeap {
	minHeap := &MinHeap{
		count:0,
		cap:cap,
	}
	
	return minHeap
}

/*
获取元素书
*/
func (this *MinIndexHeap) GetSize() int {
	return this.count
}

/*
是否空
*/
func (this *MinIndexHeap) IsEmpty() bool {
	return  this.count == 0
}

/*
是否满堆
*/
func (this *MinIndexHeap) IsFull() bool {
	return this.count == this.cap
}

/*
推进一个元素
*/
func (this *MinIndexHeap) Push(i int,v int) error {

	if(this.IsFull()){
		return errors.New("堆已满")
	}
	i += 1
	this.heap[i] = v
	this.index[this.count + 1] = i
	this.count++
	this.shiftUp(this.count)
	return nil
}

/*
获取一个元素
*/
func (this *MinIndexHeap) Pop() (int,error) {
	if(this.IsEmpty()){
		return 0,errors.New("堆已空")
	}
	var v int = this.heap[this.index[1]];
	this.index[1] = this.index[this.count]
	this.count-- 
	this.shiftDown(1)

	return v,nil
}

/*
改变值
*/
func (this *MinIndexHeap) Change(i int,v int)  {
	i += 1
	this.heap[i + 1] = v;
	for index := 1; index <= count; index++ {
		if(index == i){
			this.shiftUp(index);
			this.shiftDown(index);
			break
		}
	}
}

/*
shiftUp
*/
func (this *MinIndexHeap) shiftUp(start int) {
	var pEle int = start / 2
	if(start > 1 && this.heap[this.index[start]] < this.heap[this.index[pEle]]){
		var tem int = this.index[start]
		this.index[start] = this.index[pEle]
		this.index[pEle] = tem
		this.shiftUp(pEle)
	}
	return
}


/*
shiftDown
*/
func (this *MinIndexHeap) shiftDown(start int) {
	var j int = start * 2
	if(j <= this.count){
		if(j + 1 <= this.count && this.heap[this.index[j + 1]]  < this.heap[this.index[j]]){
			j++
		}
		if(this.heap[this.index[start]] <= this.heap[this.index[j]]){
			return
		}
		var tem int = this.index[start]
		this.index[start] = this.index[j]
		this.index[j] = tem
		this.shiftDown(j)
	}
	return
}

/*
展示堆
*/
func (this *MinIndexHeap) Show() {
	for i := 1; i <= this.count; i++ {
		fmt.Printf("%v:%d,%d \n",this.heap[i],this.heap[i * 2],this.heap[i * 2 + 1])
	}
}
