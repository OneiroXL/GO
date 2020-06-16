package main

import (
	"fmt"
)


func main(){
	fmt.Println(IntToRoman(20))
}

/*
整数转罗马数字(贪心算法)
*/
func IntToRoman(num int) string {
	// 把阿拉伯数字与罗马数字可能出现的所有情况和对应关系，放在两个数组中
	// 并且按照阿拉伯数字的大小降序排列，这是贪心选择思想
	var nums []int  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
	var romans []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

	var res string = "";

	for (num != 0){
		for index,val := range nums {
			if(num >= val){
				num = num - val
				res = res + romans[index]
                break
			}
		}
	}
	return res
}