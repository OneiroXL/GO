package main


import (
	"./list"
)

func main(){
	signList := list.SignList{}

	signList.Add(10)
	signList.Add(20)
	signList.Add(30)
	signList.Add("十")

	signList.Remove(20)
	signList.Remove("十")

	signList.Show()
}