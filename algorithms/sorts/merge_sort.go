package sort

func merge(leftArray []int32, rightArray []int32) []int32 {
	var arrSorted []int32
	leftLength := len(leftArray)
	rightLength := len(rightArray)

	l := 0
	r := 0

	for l < leftLength && r < rightLength {
		if leftArray[l] < rightArray[r] {
			arrSorted = append(arrSorted, leftArray[l])
			l += 1
		} else {
			arrSorted = append(arrSorted, rightArray[r])
			r += 1
		}
	}
	for l < leftLength {
		arrSorted = append(arrSorted, leftArray[l])
		l += 1
	}

	for r < rightLength {
		arrSorted = append(arrSorted, rightArray[r])
		r += 1
	}

	return arrSorted
}

func MergeSort(arr []int32) []int32 {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := length / 2
	leftArray := MergeSort(arr[:middle])
	rightArray := MergeSort(arr[middle:])
	return merge(leftArray, rightArray)
}
