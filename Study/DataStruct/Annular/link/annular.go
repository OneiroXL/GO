package link

import (
	"fmt"
	"../model"
)

type AnnularLink struct{
	Next *AnnularLink
	Data model.CatNode
}

func (this *AnnularLink) Add(new model.CatNode)  {
	temp := this
	if temp.Next == nil {
		temp.Data = new
		temp.Next = temp;
		return
	}
	for {
		if temp.Next == this {
			break;
		}
		temp = temp.Next;
	}

	annularLink := &AnnularLink{}
	annularLink.Data = new;
	annularLink.Next = this;

	temp.Next = annularLink;

}

func (this *AnnularLink) AddAndSort(new model.CatNode) *AnnularLink {
	temp := this;
	fmt.Println("头是:",temp)
	if temp.Next == nil {
		temp.Data = new
		temp.Next = temp;
		return this
	}
	for {
		if temp.Next.Data.No >= new.No{
			annularLink := &AnnularLink{}
			annularLink.Data = new;
			annularLink.Next = temp.Next;
			if temp.Next == this {
				this = annularLink
			}
			temp.Next = annularLink;
			return this
		} else if temp.Next == this {
			break;
		}
		temp = temp.Next;
	}

	annularLink := &AnnularLink{}
	annularLink.Data = new;
	annularLink.Next = this;

	temp.Next = annularLink;
	return this
}

func (this *AnnularLink) Delete(no int) *AnnularLink {
	temp := this
	if(temp.Next == this){
		this.Next = nil
		return this
	}
	for {
		if(temp.Next.Data.No == no){
			if(temp.Next == this){
				this = temp
			}
			temp.Next = temp.Next.Next
			return this
		}
		temp = temp.Next;
	}
} 

func (this *AnnularLink) Show(){
	temp := this;
	if temp.Next == nil {
		fmt.Printf("空空如也 \n");
		return
	}
	for {
		fmt.Printf("我是%v,喵喵喵 \n",temp.Data.Name);
		if temp.Next == this {
			break;
		}
		temp = temp.Next;
	}
}