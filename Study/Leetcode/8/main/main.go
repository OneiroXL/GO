package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(Atio("9223372036854775808"))
	fmt.Println(Test("9223372036854775808"))
	//fmt.Println((9223372036854775800 + 8))
}

// Atio string to int
func Atio(str string) int {
	//remove ' '
	var tem string = strings.TrimSpace(str)
	if len(tem) == 0 {
		return 0
	}
	// op
	var op int = 1
	//string to int
	var firstChar int = int(tem[0])
	if (firstChar == '-' || firstChar == '+') || (firstChar <= '9' && firstChar >= '0') {
		if firstChar == '-' {
			op = -1
		}
		var ans int = 0
		if firstChar <= '9' && firstChar >= '0' {
			ans = ans*10 + firstChar - '0'
		}
		for i := 1; i < len(tem); i++ {
			if tem[i] >= '0' && tem[i] <= '9' {
				var digit int = int(tem[i]) - '0'
				ans = ans*10 + digit
			} else {
				return ans * op
			}
			// max value testing
			if ans > math.MaxInt32 {
				if op == -1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		}
		return ans * op
	} else {
		return 0
	}
}

func Test(str string) int {
	//去掉收尾空格
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 数值最大检测
		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result
}
