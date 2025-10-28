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
	getleft(idx int) int
	getright(idx int) int
	getParenint(idx int) (int, error)
	getlen() int
}

type HeapImpl struct {
	items []int
}

func NewHeap() HeapImpl {
	return HeapImpl{}
}

func (h *HeapImpl) getParent(idx int) (int, error) {
	if idx == 0 {
		return 0, errors.New("root cannot have parent")
	}
	return idx / 2, nil
}

func (h *HeapImpl) bubbleup(idx int)  {
	paridx, err := h.getParent(idx)
	if err != nil{
		return
	}
	if h.items[paridx] > h.items[idx]{
		h.items[paridx],h.items[idx] = h.items[idx], h.items[paridx]
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