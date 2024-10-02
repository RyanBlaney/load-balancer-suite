package priorityqueue

import (
	"container/heap"
	"errors"
)

// PriorityQueue
//
//	manages a priority queue with heap functionality
type PriorityQueue struct {
	itemHeap itemHeap
	lookup   map[interface{}]*item
}

// item
//
//	represents an element in the priority queue
type item struct {
	value    interface{}
	priority float64
	index    int
}

// NewPriorityQueue
//
//	creates a new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	h := &itemHeap{}
	heap.Init(h)
	return &PriorityQueue{
		itemHeap: *h,
		lookup:   make(map[interface{}]*item),
	}
}

// Len
//
//	returns the length of the priority queue
func (p *PriorityQueue) Len() int {
	return len(p.itemHeap)
}

// Insert
//
//	inserts a value with a given priority into the priority queue
func (p *PriorityQueue) Insert(v interface{}, priority float64) {
	_, ok := p.lookup[v]
	if ok {
		return // Item already exists, don't insert it again
	}
	newItem := &item{
		value:    v,
		priority: priority,
	}
	heap.Push(&p.itemHeap, newItem)
	p.lookup[v] = newItem
}

// Pop
//
//	removes and returns the highest priority item from the queue
func (p *PriorityQueue) Pop() (interface{}, error) {
	if len(p.itemHeap) == 0 {
		return nil, errors.New("empty queue")
	}
	item := heap.Pop(&p.itemHeap).(*item)
	delete(p.lookup, item.value)
	return item.value, nil
}

// UpdatePriority
//
//	updates the priority of a given item in the queue
func (p *PriorityQueue) UpdatePriority(v interface{}, newPriority float64) {
	item, ok := p.lookup[v]
	if !ok {
		return
	}
	item.priority = newPriority
	heap.Fix(&p.itemHeap, item.index)
}

// itemHeap
//
//	implements heap.Interface for a slice of items
type itemHeap []*item

// Len
//
//	is the number of elements in the collection
func (ih itemHeap) Len() int {
	return len(ih)
}

// Less
//
//	reports whether the element with index i should sort before the element with index j
func (ih itemHeap) Less(i, j int) bool {
	return ih[i].priority < ih[j].priority
}

// Swap
//
//	swaps the elements with indexes i and j
func (ih itemHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
	ih[i].index = i
	ih[j].index = j
}

// Push
//
//	adds an item to the heap
func (ih *itemHeap) Push(x interface{}) {
	it := x.(*item)
	it.index = len(*ih)
	*ih = append(*ih, it)
}

// Pop
//
//	removes and returns the last element from the heap
func (ih *itemHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	it := old[n-1]
	*ih = old[0 : n-1]
	return it
}
