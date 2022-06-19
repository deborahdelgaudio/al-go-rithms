package arrays

func BubbleSort(arr []int) []int {
	n := len(arr);
	for j := n; j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if arr[i] > arr[i+1]{
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr;
}

func InsertSort(arr []int) []int {
	len := len(arr)
    for i:=1; i < len; i++{
		i_elem := arr[i]
		j := i-1 
		for j>=0 && arr[j] > i_elem {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = i_elem
	}
	return arr
}

func merge(left []int, right []int) (result []int) {
	result = make([]int, len(left) + len(right))
	a := 0
	b := 0
	c := 0
	for a < len(left) && b < len(right) {
		if left[a] < right[b]{
			result[c] = left[a]
			a ++
		} else {
			result[c] = right[b]
			b ++
		}
		c ++
	}
	for a < len(left){
		result[c] = left[a]
		a ++
		c ++
	}
	for b < len(right) {
		result[c] = right[b]
		b ++
		c ++
	}
	return result
}

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	// split recursively
	middlepoint := int(len(list) / 2)
	left := make([]int, middlepoint)
	right := make([]int, len(list)-middlepoint)
	for i := 0; i < len(list); i++ {
		if i < middlepoint{
			left[i] = list[i]
		}else{
			right[i-middlepoint] = list[i]
		}
	}
	return merge(MergeSort(left), MergeSort(right))
}