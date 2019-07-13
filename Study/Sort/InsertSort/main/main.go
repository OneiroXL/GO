package main

import(
	"fmt"
	"../sort"
)

func main()  {

	var arr [5]int = [5]int{1,8,9,6,3}

	insertSort := sort.NewInsertSort();


	fmt.Println(insertSort.Sort(arr[:]))

}