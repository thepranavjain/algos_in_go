package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	arrays := [][]int{{90, 70, 50, 80, 60, 85}, {9, 7, 5, -3, 4, 2, 1, -55}, {9, 8, 7, 6, 5, 4, 3}, {0, 1, 2, 3, 4, 5}, {3, 2}, {2, 3}}
	for _, ar := range arrays {
		ar = BubbleSort(ar)
		for i := 1; i < len(ar); i++ {
			if ar[i] < ar[i-1] {
				t.Error("Array", ar, "not sorted")
			}
		}
	}
}
