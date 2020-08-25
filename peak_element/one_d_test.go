package peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	arrs := [][]int{{1,2,3,1},{1,2,1,3,5,6,4},{2},{6,4},{2,4},{2,4,6},{6,4,2,0},{2,6,4},{1,1,1,1,1}}
	ans := [][]int{{2},{1,5},{0},{0},{1},{2},{0},{1},{0,1,2,3,4}} // <-- Possible indices of peak elements
	for i, arr := range arrs {
		res := FindPeakElement(arr)
		correct := false
		for _, v := range ans[i] {
			if v == res {
				correct = true
			}
		}
		if !correct {
			t.Errorf("Error in FindPeakElement() on i/p %v. Allowed answers %v. Received %v", arr, ans[i], res)
		}
	}
}