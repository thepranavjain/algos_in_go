package searching

func BinarySearch(ar []int, val int) int {
	n := len(ar)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if ar[0] == val {
			return 0
		} else {
			return -1
		}
	}
	low := 0
	high := n - 1
	mid := (high + low) / 2
	if ar[mid] == val {
		return mid
	}
	if ar[mid] < val {
		subArrayIndex := BinarySearch(ar[mid+1:n], val)
		if subArrayIndex == -1 {
			return -1
		}
		return mid + 1 + subArrayIndex
	}
	if ar[mid] > val && mid != low {
		return BinarySearch(ar[low:mid], val)
	}
	return -1
}
