package heap

// PriorityQueue will return maximum values on popping
type PriorityQueue struct {
	h *Heap
}

// NewPriorityQueue creates an empty object and returns a ref to it
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		h: New(),
	}
}

// Push the values into the PriorityQueue
// Complexity O(log(n))
func (pq *PriorityQueue) Push(val int) {
	pq.h.Push(-val)
}

// Top returns the maximum value without removing it
// Complexity O(1)
func (pq *PriorityQueue) Top() int {
	return -pq.h.Top()
}

// Pop returns the maximum value while removing it from PriorityQueue
// Complexity O(log(n))
func (pq *PriorityQueue) Pop() int {
	return -pq.h.Pop()
}
