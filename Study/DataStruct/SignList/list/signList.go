package list

import(
	"fmt"
)

type SignList struct {
	Next *SignList
	Data interface{}
}

/*
添加
*/
func (this *SignList) Add(data interface{}) {
	temp := this
	for {
		if(temp.Next == nil){
			signList := &SignList{
				Data:data,
			}
			temp.Next = signList 
			return
		}
		temp = temp.Next
	}
}

/*
移除链表
*/
func (this *SignList) Remove(data interface{}){
	temp := this
	for {
		if(temp.Next.Data == data){
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}
}

/*
打印链表
*/
func (this *SignList) Show(){
	temp := this
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}
