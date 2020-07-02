package main

import (
	"fmt"
)

func main() {
 	fmt.Println(IsPalindrome(-111121111))
}


func IsPalindrome(val int) bool {
	op := 1
	if(val < 0){
		op = -1
	}
	tem := val
	res := 0
	for(tem != 0){
		var rem int = tem % 10
		res = res * 10 + rem
		tem = tem / 10
	}
	if (res * op == val) {
		return true
	} else{
		return false
	}
} 
