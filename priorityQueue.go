package priorityqueue

import (
	"container/heap"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	data    []interface{}
	compare func(item1, item2 interface{}) bool
}

// New creates new priority queue.
// compare function should get 2 elements of queue and
// return true if the first argument must be early in queue then the second one.
func New(compare func(item1, item2 interface{}) bool) *PriorityQueue {
	pq := PriorityQueue{}
	pq.compare = compare
	return &pq
}

// AddItem adds an item to the priority queue. USE IT INSTEAD OF Push
func (pq *PriorityQueue) AddItem(x any) {
	heap.Push(pq, x)
}

// GetItem pops an item from the priority queue. USE IT INSTEAD OF Pop
func (pq *PriorityQueue) GetItem() interface{} {
	return heap.Pop(pq)
}

func (pq *PriorityQueue) Len() int {
	return len(pq.data)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.compare(pq.data[i], pq.data[j])
}

// Push SHOULDN'T BE USED MANUALLY, use AddItem instead.
// Push is used by container/heap functions.
func (pq *PriorityQueue) Push(x any) {
	pq.data = append(pq.data, x)
}

// Pop SHOULDN'T BE USED MANUALLY, use GetItem instead.
// Pop is used by container/heap functions.
func (pq *PriorityQueue) Pop() any {
	n := pq.Len()
	item := pq.data[n-1]
	pq.data[n-1] = nil // avoid memory leak
	pq.data = pq.data[0 : n-1]
	return item
}
