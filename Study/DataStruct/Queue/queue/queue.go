package queue

import (
	"errors"
	"fmt"
)

var (
	rear      int
	front     int
	count     int
	max       int
	queueData []interface{}
)

type Queue struct {
}

func NewQueue(cap int) *Queue {
	rear = 0
	front = 0
	count = 0
	max = cap
	queueData = make([]interface{}, cap)
	return &Queue{}
}

/*
入队
*/
func (this *Queue) Push(v interface{}) error {
	if this.IsFull() {
		return errors.New("Queue Is Full")
	}
	if front == max {
		front = 0
	}
	queueData[front] = v
	front++
	count++
	return nil
}

/*
出队
*/
func (this *Queue) Pop() (interface{}, error) {
	if this.IsEmpty() {
		return nil, errors.New("Queue Is Empty")
	}
	if rear == max {
		rear = 0
	}
	v := queueData[rear]
	rear++
	count--
	return v, nil
}

/*
是否满对
*/
func (this *Queue) IsFull() bool {
	return len(queueData) == count
}

/*
是否空对
*/
func (this *Queue) IsEmpty() bool {
	return count == 0
}

/*
打印队列
*/
func (this *Queue) Show() {
	temp := count
	tempRear := rear
	for temp > 0 {
		if tempRear >= max {
			tempRear = 0
		}
		fmt.Printf("%d->", queueData[tempRear])
		tempRear++
		temp--
	}
	fmt.Println()
}
