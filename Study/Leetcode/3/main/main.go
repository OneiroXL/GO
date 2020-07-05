package main

import (
	"fmt"
)

func main() {
	fmt.Println(LengthOfLongestSubstring("abba"))
}


func LengthOfLongestSubstring(s string) int {
	var maxLength int = 0
	var flag int = 0
	var mapIndex map[byte]int = make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		if (MapExistKey(mapIndex,s[i])) { //如果map中存在当前字符则维护flag
			if(mapIndex[s[i]] + 1 > flag ){//此处做判断是因为要是当前的字符之前出现的位置已经排除在外了（即mapIndex[s[i]] + 1 < flag）
				flag = mapIndex[s[i]] + 1//这里的加一是因为mapIndex[s[i]]是当前字符本身 加一后要把它也排除在外
			}
		}
		mapIndex[s[i]] = i
		var length int = i - flag + 1 //这里的加一是因为 若s=zx s[0] = 'z' s[1] = 'x'  s的长度是1-0+1而不是1 -0   
		if(length > maxLength){
			maxLength = length
		}
	}
	return maxLength
}


func MapExistKey(mapIndex map[byte]int,key byte ) bool {
	if _, ok := mapIndex[key]; ok {
		return true	
	}
	return false
}