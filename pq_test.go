package priorityqueue

import (
	`container/heap`
	`fmt`
	`testing`
)

type item struct {
	val 		int
	priority	int
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue(func(item1, item2 interface{}) bool {
		return item1.(item).priority < item2.(item).priority
	})
	for i := 0; i < 10; i++ {
		heap.Push(pq, item{i, i})
	}
	for i := 9; i >= 0; i-- {
		heap.Push(pq, item{i, i})
	}
	for pq.Len() > 0 {
		item := heap.Pop(pq).(item)
		fmt.Println(item.priority, item.val)
	}
}
