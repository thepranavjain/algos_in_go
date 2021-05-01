package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	arrays := [][]int{
		{30, 40, 50, 70, 85, 90, 100},
		{30, 40, 50, 70, 85, 90, 100},
		{30, 40, 50, 70, 85, 90, 100},
		{},
		{1},
		{2},
		{3},
		{1, 5},
		{1, 6},
		{3, 7},
		{1, 2, 3},
		{1, 3, 5},
		{1, 5, 8},
		{4, 6, 9},
	}
	toSearch := []int{40, 90, 109, 997, 1, 5, 1, 5, 1, 8, 1, 3, 8, 1}
	expectedResults := []int{1, 5, -1, -1, 0, -1, -1, 1, 0, -1, 0, 1, 2, -1}
	for i, ar := range arrays {
		result := BinarySearch(ar, toSearch[i])
		if result != expectedResults[i] {
			t.Error(
				"Searching",
				toSearch[i],
				"in",
				ar,
				"Expected result:",
				expectedResults[i],
				"Received result:",
				result,
			)
		}
	}
}
