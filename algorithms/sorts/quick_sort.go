package sort

func QuickSort(arr []int32, start, end int) []int32 {
	if start < end {
		var part int
		arr, part = partition(arr, start, end)
		arr = QuickSort(arr, start, part-1)
		arr = QuickSort(arr, part+1, end)
	}
	return arr
}

func partition(arr []int32, start, end int) ([]int32, int) {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return arr, i
}
