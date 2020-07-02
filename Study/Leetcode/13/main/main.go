package main

import (
	"fmt"
)

func main()  {
	
	fmt.Println(RomanToInt("LIV"))
}

/*
主要注意点
如果上一个元素大于或等于本次的元素则相加
如果上一个元素小于本次的元素则相减
*/

func RomanToInt(s string) int {
	romanToIntMap := map[byte]int {
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var num int = 0;
	temStr := s
	temNum := 0
	for i := len(temStr) - 1; i >= 0; i-- {
		val := romanToIntMap[temStr[i]]
		if(val >= temNum){
			num = num +  val
		} else{
			num = num - val
		}
		temNum = val
	}
	return num
}