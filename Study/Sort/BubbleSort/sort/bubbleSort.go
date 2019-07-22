package sort


import (

)


type BubbleSort struct{

}


func NewBubblingSort() *BubbleSort {
	return &BubbleSort{}
}

func (this *BubbleSort) Sort(v []int) []int {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v) - i - 1; j++ {
			if(v[j] > v[j + 1]){
				temp := v[j]
				v[j] = v[j + 1]
				v[j + 1] = temp
			}
		}
	}
	return v
}