
package main

import (
	"fmt"
	"sort"
)

/*
题目要求:给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main(){
	var numArr []int  = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(ThreeSum(numArr))
}

/*
考察三指针
*/
func ThreeSum(nums []int) [][]int {
	//结果
	var result [][]int = [][]int{}
	//排序(从小到大)
	sort.Ints(nums)

	var k int = 0
	var i int = 0
	var j int = len(nums) - 1
	
	for(k < len(nums) - 2){
		if(nums[k] > 0){//这里是因为数组是从小到大排序，所以当K大于0 说明后面的数值都大于0所以跳过就好
			break;
		}
		if (k > 0 && nums[k] == nums[k - 1]){//这里是因为nums[k]等于nums[k - 1]的话 说明nums[k - 1]这个值已经被考察过了，所以直接跳过就可以了
			k++
			continue
		}
		i =  k + 1 //从K+1这个地方开始考察(相当于第二个元素)
		j =  len(nums) - 1//考察到len(nums)-1这个位置结束(相当于第三个元素)
		for(i < j){//当i和j不相遇时 说明中间还以后没有考察完的元素
			s := nums[k] + nums[i] + nums[j] //此时应为第一个元素nums[k],第二个元素nums[i],第三个元素nums[j]都已经定下 ，所以算出三个元素的结果
			if (s < 0){
				i += 1//如果结果小于0 则从左边考察下一个元素
                for (i < j && nums[i] == nums[i - 1]){ // 这里算是一个小优化 要是下一个元素等于上一个考察元素 也没有必要去考察了 所以继续下一个元素(因为不能存在重复的)
					i += 1
				} 
			} else if (s > 0){
				j -= 1 //如果结果大于0 则从右边考察下一个元素
				for(i < j && nums[j] == nums[j + 1]){// 这里算是一个小优化 要是下一个元素等于上一个考察元素 也没有必要去考察了 所以继续下一个元素(因为不能存在重复的)
					j -= 1
				} 
			} else{
				result = append(result,[]int{nums[k] ,nums[i],nums[j]})//当不属于上面两种情况 结果一定等于0 这时候维护result
				i += 1
				j -= 1
				for (i < j && nums[i] == nums[i - 1]){
					i += 1
				} 
				for (i < j && nums[j] == nums[j + 1]){
					j -= 1
				} 
			}
		}
		k++
	}
	return result
}