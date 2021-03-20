package heap

import (
	"testing"
)

func TestNew(t *testing.T) {
	h := New()
	if len(h.arr) != 0 {
		t.Error("Error in heap.New(). Empty array not initialized.")
	}
}

func TestPush(t *testing.T) {
	testCases := [][]int{
		{35, 33, 42, 10, 14, 19, 27, 44, 26, 31},
		{6, 10, 2, 4, 1},
		{5},
		{5, 6},
		{6, 5},
		{3, 2, 1},
	}
	expectedResults := [][]int{
		{10, 14, 19, 26, 31, 42, 27, 44, 35, 33},
		{1, 2, 6, 10, 4},
		{5},
		{5, 6},
		{5, 6},
		{1, 3, 2},
	}
	for index, sampleArr := range testCases {
		expectedArr := expectedResults[index]
		h := New()
		for _, v := range sampleArr {
			h.Push(v)
		}
		if len(h.arr) != len(expectedArr) {
			t.Errorf("heap.Push() failed. Length of heap's array. Expected:%d. Received:%d", len(sampleArr), len(h.arr))
		}
		for i := range expectedArr {
			if h.arr[i] != expectedArr[i] {
				t.Errorf("heap.Push() falied. At index %d expected %d. Received %d", i, expectedArr[i], h.arr[i])
			}
		}
	}
}

func TestPop(t *testing.T) {
	// testCases include pre-determined heap arrays
	testCases := [][]int{
		{10, 14, 19, 26, 31, 42, 27, 44, 35, 33},
		{1, 2, 6, 10, 4},
		{5},
		{5, 6},
		{1, 3, 2},
	}
	// expected order of popping of integers
	expectedResults := [][]int{
		{10, 14, 19, 26, 27, 31, 33, 35, 42, 44},
		{1, 2, 4, 6, 10},
		{5},
		{5, 6},
		{1, 2, 3},
	}
	heaps := []*Heap{}
	for _, v := range testCases {
		heaps = append(heaps, &Heap{arr: v})
	}
	for i, expectedRes := range expectedResults {
		currentHeap := heaps[i]
		for j, expectedNum := range expectedRes {
			receivedNum := currentHeap.Pop()
			if receivedNum != expectedNum {
				t.Errorf("heap.Pop() failed. %d call on heap: %v. Expected: %d. Received %d",
					j, testCases[i], expectedNum, receivedNum,
				)
			}
			if len(currentHeap.arr) != len(expectedRes)-j-1 {
				t.Errorf("heap.Pop() failed. Expected length after %d pops: %d. Found: %d",
					j, len(expectedRes)-j-1, len(currentHeap.arr),
				)
			}
		}
	}
}
