package main

import (
	"fmt"
	"math"
)


func main()  {
	var intArr []int =  []int{7,1,1,1,1,1,10,8}
	fmt.Println(MaxAreaTwo(intArr[:]))
}


func MaxArea(height []int) int {
	var area int = 0
	for keyOne, valOne := range height {
		for keyTwo, valTwo := range height {
			var w int  = int(math.Abs(float64(keyOne - keyTwo)))
			var h int  = 0;
			if(valOne > valTwo){
				h = valTwo
			} else {
				h = valOne
			}
			res := w*h
			if(res > area){
				area = res
			}
		}
	}
	return area
}

func MaxAreaTwo(height []int) int {
	var area int = 0
	var i int = 0
	var j int = len(height) - 1
	for range height {
		res := 0
		if(height[i] < height[j]){
			res = height[i] * int(math.Abs(float64(j - i)))
			i++
		} else {
			res = height[j] * int(math.Abs(float64(j - i)))
			j--
		}
		if(i == j){
			break
		}
		if(res > area){
			area = res
		}
	}
	return area
}