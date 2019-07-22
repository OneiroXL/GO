package sort

type SelectSort struct{

}
/*选择排序的思想
1.从第一个位置开始找，从当前位置+1到结束位置中找到最小的"值"的位置
2.将每次寻找的起始位置与找到的最小值的位置的值进行交换
*/

func NewSelectSort() *SelectSort {
	return &SelectSort{};
}

func (this *SelectSort) Sort(v []int) []int {
	for i := 0; i < len(v); i++ {
		var minInsdex int = i;
		for j := i + 1; j < len(v); j++ {
			if(v[minInsdex] > v[j]){
				minInsdex = j
			}
		}
		temp := v[i];
		v[i] = v[minInsdex];
		v[minInsdex] = temp;
	}

	return v;
}