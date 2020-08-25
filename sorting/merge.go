package sorting

func MergeSort(ar []int) []int {
	n := len(ar)
	if n == 1 {
		return ar
	}

	// Divide -------------->
	lArr := MergeSort(ar[0:n/2])
	rArr := MergeSort(ar[n/2:n])

	var r,l int
	var res []int
	// Merge ----------------->
	for r < len(rArr) && l < len(lArr) {
		if rArr[r] > lArr[l] {
			res = append(res, lArr[l])
			l++
			continue
		}
		res = append(res, rArr[r])
		r++
	}
	// Only one of the below loops will run and it will add all remaining elements to res
	for r < len(rArr) {
		res = append(res,rArr[r])
		r++
	}
	for l < len(lArr) {
		res = append(res, lArr[l])
		l++
	}
	return res
}
