package main

import (
	_"fmt"
	"../link"
	"../model"
)

func main()  {
	annularLink := &link.AnnularLink{}

	catNodeOne := model.CatNode{
		No:1,
		Name:"小白",
	} 
	catNodeTwo := model.CatNode{
		No:2,
		Name:"小黑",
	}
	catNodeThree := model.CatNode{
		No:3,
		Name:"小花",
	}
	catNodeFour := model.CatNode{
		No:4,
		Name:"小萌",
	}
	catNodeFive := model.CatNode{
		No:5,
		Name:"小幸",
	}

	// annularLink.Add(catNodeOne);
	// annularLink.Add(catNodeTwo);
	// annularLink.Add(catNodeThree);
	// annularLink = annularLink.Delete(1);


	annularLink = annularLink.AddAndSort(catNodeOne);
	annularLink = annularLink.AddAndSort(catNodeFour);
	annularLink = annularLink.AddAndSort(catNodeThree);
	annularLink = annularLink.AddAndSort(catNodeFive);
	annularLink = annularLink.AddAndSort(catNodeTwo);



	annularLink.Show()


}