package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	arrays := [][]int{{9,7,5,-3,4,2,1,-55},{9,8,7,6,5,4,3},{0,1,2,3,4,5},{3,2},{2,3}}
	for _, ar := range arrays {
		ar = MergeSort(ar)
		for i := 1; i < len(ar); i++ {
			if ar[i] < ar[i-1] {
				t.Error("Array",ar,"not sorted")
			}
		}
	}
}

