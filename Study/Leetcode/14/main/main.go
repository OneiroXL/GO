package main

import (
	"fmt"
)


func main()  {
	var strArr [2]string = [2]string{"abdf","abdfasdasdasd"}

	fmt.Println(LongestCommonPrefix(strArr[:]))
}

func LongestCommonPrefix(strs []string) string {
	var res string = ""
	if(len(strs) == 0){
		return res
	}
	var lenght int = len(strs[0]) 
	for i := 0; i < len(strs); i++ {
		if(len(strs[i]) < lenght){
			lenght = len(strs[i])
		} 
	}
	var i int = 0
	for(i < lenght){
		var flag bool = true
		var pre byte = strs[0][i]
		for _, val := range strs {
			if(val[i] != pre){
				flag = false
			}
		}
		if(flag){
			res  = res + string(pre)
		} else {
			return res
		}
		i++
	}

	return res
}

