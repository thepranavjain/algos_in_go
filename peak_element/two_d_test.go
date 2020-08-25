package peak_element

import "testing"

func TestFindHillTop(t *testing.T) {
	inputs := [][][]int{{{10,8,10,10},{14,13,12,11},{15,9,11,21},{16,17,19,20}},
						{{10,20,15},{21,30,14},{7,16,32}},
						{{4,4,4},{4,4,4},{4,4,4}},
						{{2}},
						{{2,1}},
						{{1,2},{3,4}},
						{{1,2},{3,6},{5,4}},
						{{1,2,3},{6,4,5}}}
	answers := [][][]int{{{2,3}},
						{{1,1},{2,2}},
						{{0,0},{0,1},{0,2},{1,0},{1,1},{1,2},{2,0},{2,1},{2,2}},
						{{0,0}},
						{{0,0}},
						{{1,1}},
						{{1,1},{2,0}},
						{{1,0},{1,2}}}
	for i, arr := range inputs {
		x, y := FindHillTop(arr)
		correct := false
		for _, ans := range answers[i] {
			if x == ans[0] && y == ans[1] {
				correct = true
			}
		}
		if !correct {
			t.Errorf("Error in FindHillTop() on i/p %v. Allowed answers %v. Received %v", arr, answers[i], []int{x,y})
		}
	}
}
