package sorting

func InsertionSort(ar []int) []int {
	// [0,1,2,3,4,5]
	n := len(ar)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if ar[j] < ar[j-1] {
				tmp := ar[j]
				ar[j] = ar[j-1]
				ar[j-1] = tmp
			}
		}
	}
	return ar
}
