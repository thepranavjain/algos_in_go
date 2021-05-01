package sorting

func BubbleSort(ar []int) []int {
	n := len(ar)
	for {
		swaps := 0
		for i := 0; i < n-1; i++ {
			if ar[i] > ar[i+1] {
				ar[i], ar[i+1] = ar[i+1], ar[i]
				swaps++
			}
		}
		if swaps == 0 {
			break
		}
	}
	return ar
}
