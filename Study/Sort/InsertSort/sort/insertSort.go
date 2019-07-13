package sort

type InsertSort struct{

}

/*插入排序思想
1.便利整个元素
2.将每个元素和它之前的所有元素相比较
3.如果小于它，则此时让与它比较的元素的向后移
4.如果大于它，则此时结束比较
5.每次比较结束时，将它与它最后进行比较的元素进行交换
*/

func NewInsertSort() *InsertSort {
	var insertSort *InsertSort = &InsertSort{}
	return insertSort
}

func (this *InsertSort) Sort(v []int) []int {
	for i := 1; i < len(v) ;i++ {
		var j int = 0
		var k int = v[i]
		for j = i; j > 0 && v[j - 1] > k; j-- {
			v[j] = v[j - 1]
		}
		v[j] = k
	}

	return v
}