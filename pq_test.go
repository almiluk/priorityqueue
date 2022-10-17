package priorityqueue

import (
	"fmt"
	"testing"
)

type item struct {
	val      int
	priority int
}

func TestPriorityQueue(t *testing.T) {
	pq := New(func(item1, item2 interface{}) bool {
		return item1.(item).priority < item2.(item).priority
	})
	for i := 0; i < 10; i++ {
		pq.AddItem(item{i, i})
	}
	for i := 9; i >= 0; i-- {
		pq.AddItem(item{i, i})
	}
	for pq.Len() > 0 {
		item := pq.GetItem().(item)
		fmt.Println(item.priority, item.val)
	}
}
