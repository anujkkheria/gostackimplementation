package main

import (
	"errors"
	"fmt"
)

type Heap interface {
	Pop() (int, error)
	Push(value int)
	Peek() int
	bubbleup(i int)
	heapify()
	getleft() int
	getright() int
	getParent() (int, error)
	getlen() int
}

type HeapImpl struct {
	items []int
}

func NewHeap() HeapImpl {
	return HeapImpl{}
}

func (h *HeapImpl) getParent(ind int) (int, error) {
	if ind == 0 {
		return 0, errors.New("root cannot have parent")
	}
	return ind / 2, nil
}

func (h *HeapImpl) bubbleup(ind int)  {
	paridx, err := h.getParent(ind)
	if err != nil{
		return
	}
	if h.items[paridx] > h.items[ind]{
		h.items[paridx],h.items[ind] = h.items[ind], h.items[paridx]
		h.bubbleup(paridx)
	}
}

func (h *HeapImpl) Push(val int) {
	h.items = append(h.items, val)
	i := len(h.items) - 1
	h.bubbleup(i)
}

func main() {

	heap := NewHeap()

	heap.Push(2)
	heap.Push(3)
	heap.Push(1)

	fmt.Println(heap)
}