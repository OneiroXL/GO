package help

import (
	"errors"
	"fmt"
	"../../Model"
)

type MinHeap struct {
	heap [100]*model.Edge
	count int
	cap int
}

/*
创建一个最小堆
*/
func NewMinHeap(cap int) *MinHeap {
	minHeap := &MinHeap{
		count:0,
		cap:cap,
	}
	
	return minHeap
}

/*
获取元素书
*/
func (this *MinHeap) GetSize() int {
	return this.count
}

/*
是否空
*/
func (this *MinHeap) IsEmpty() bool {
	return  this.count == 0
}

/*
是否满堆
*/
func (this *MinHeap) IsFull() bool {
	return this.count == this.cap
}

/*
推进一个元素
*/
func (this *MinHeap) Push(v *model.Edge) error {

	if(this.IsFull()){
		return errors.New("堆已满")
	}
	this.heap[this.count + 1] = v
	this.count++
	this.shiftUp(this.count)
	return nil
}

/*
获取一个元素
*/
func (this *MinHeap) Pop() (*model.Edge,error) {
	if(this.IsEmpty()){
		return nil,errors.New("堆已空")
	}
	var v *model.Edge = this.heap[1];
	this.heap[1] = this.heap[this.count]
	this.count-- 
	this.shiftDown(1)

	return v,nil
}

/*
shiftUp
*/
func (this *MinHeap) shiftUp(start int) {
	var pEle int = start / 2
	if(start > 1 && this.heap[start].Weight.(int) < this.heap[pEle].Weight.(int)){
		var tem *model.Edge = this.heap[start]
		this.heap[start] = this.heap[pEle]
		this.heap[pEle] = tem
		this.shiftUp(pEle)
	}
	return
}


/*
shiftUp
*/
func (this *MinHeap) shiftDown(start int) {
	var j int = start * 2
	if(j <= this.count){
		if(j + 1 <= this.count && this.heap[j + 1].Weight.(int)  < this.heap[j].Weight.(int)){
			j++
		}
		if(this.heap[start].Weight.(int) <= this.heap[j].Weight.(int)){
			return
		}
		var tem *model.Edge = this.heap[start]
		this.heap[start] = this.heap[j]
		this.heap[j] = tem
		this.shiftDown(j)
	}
	return
}

/*
展示堆
*/
func (this *MinHeap) Show() {
	for i := 1; i <= this.count; i++ {
		fmt.Printf("%v:%d,%d \n",this.heap[i].Weight,this.heap[i * 2].Weight,this.heap[i * 2 + 1].Weight)
	}
}
