package main

import (
	"fmt"
	"strings"
)


func main()  {
	var strArr [3]string = [3]string{"flower","flow","flight"}

	fmt.Println(LongestCommonPrefix(strArr[:]))
}

func LongestCommonPrefix(strs []string) string {
	var temStr string = ""

	for _, val := range strs {
		temStr = temStr + val + ","
	}

	temStr = strings.Trim(temStr,",")
	fmt.Println(temStr)
	
	var flagIndex int = 0
	var length int = 0
	var res int = 0
	for i := 1; i < len(temStr); i++ {
		if(temStr[i] == temStr[flagIndex]){
			flagIndex++
		}
		if(length > res){
			res = length	
		}
		length = i - flagIndex
	}



	return ""
}