package heap

// Heap is the implementation of a min-heap
type Heap struct {
	arr []int
}

// New creates an empty heap object and returns the pointer to it
func New() *Heap {
	return &Heap{
		arr: []int{},
	}
}

// Push an element into the min heap
// Complexity O(log(n))
func (h *Heap) Push(num int) {
	h.arr = append(h.arr, num)
	newElemIndex := len(h.arr) - 1
	parent := (newElemIndex - 1) / 2
	for h.arr[parent] > h.arr[newElemIndex] {
		// swap elements in h.arr
		temp := h.arr[parent]
		h.arr[parent] = h.arr[newElemIndex]
		h.arr[newElemIndex] = temp
		// update pointer to the new elem
		newElemIndex = parent
		// Check for new parent
		parent = (newElemIndex - 1) / 2
		if parent < 0 {
			break
		}
	}
}

// Top returns the root. Doesn't pop it off the heap.
// Complexity O(1)
func (h *Heap) Top() int {
	return h.arr[0]
}

// Pop removes the minimum (root) value from the heap and returns it
// Complexity O(log(n))
func (h *Heap) Pop() int {
	res := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	index := 0
	indexOfSmallest := h.findSmallestInSelfAndChildren(index)
	for index != indexOfSmallest {
		temp := h.arr[indexOfSmallest]
		h.arr[indexOfSmallest] = h.arr[index]
		h.arr[index] = temp
		index = indexOfSmallest
		indexOfSmallest = h.findSmallestInSelfAndChildren(index)
	}
	return res
}

// Returns the index of the smallest number among
// the number at index and its children (if they exist)
func (h *Heap) findSmallestInSelfAndChildren(index int) int {
	indexOfSmallest := index
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2
	if leftChildIndex < len(h.arr) && h.arr[leftChildIndex] < h.arr[indexOfSmallest] {
		indexOfSmallest = leftChildIndex
	}
	if rightChildIndex < len(h.arr) && h.arr[rightChildIndex] < h.arr[indexOfSmallest] {
		indexOfSmallest = rightChildIndex
	}
	return indexOfSmallest
}
