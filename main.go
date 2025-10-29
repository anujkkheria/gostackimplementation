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
	heapify(idx int)
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

func (h *HeapImpl) getleft (idx int)(int,error){
	heaplength := len(h.items)
	leftidx := 2*idx+1
	if heaplength <= leftidx{
		return 0 , errors.New("index out of bounds")
}
 return leftidx ,nil	
}

func (h *HeapImpl) getright(idx int)(int, error){
	heaplength := len(h.items)
	rightidx := 2*idx+2
	if heaplength <= rightidx{
		return 0 , errors.New("index out of bounds")
}
 return rightidx ,nil	
}

	func (h *HeapImpl) heapify(idx int){
		smallestIdx := idx
		leftidx,err :=  h.getleft(idx)
		if err != nil{
			return
		}
		if h.items[leftidx] < h.items[smallestIdx]{
			smallestIdx = leftidx
		}
	rightidx,err:=  h.getright(idx)
if err == nil{
	if h.items[rightidx] < h.items[smallestIdx]{
		smallestIdx = rightidx
	}
}
h.items[smallestIdx], h.items[idx] = h.items[idx],h.items[smallestIdx]
h.heapify(smallestIdx)
	}

func (h *HeapImpl) Push(val int) {
	h.items = append(h.items, val)
	i := len(h.items) - 1
	h.bubbleup(i)
}

func (h *HeapImpl) Pop()(int,error){
heaplen :=  h.getlen()
if heaplen < 1{
	return 0, errors.New("heap is empty")
}
poppedItem := h.items[0]

if heaplen > 1 {
	h.items[0], h.items[heaplen-1] = h.items[heaplen-1],h.items[0]
	h.items = h.items[:heaplen-1]
	h.heapify(0)
}else{
	h.items = h.items[:heaplen-1]
}

return poppedItem, nil 
}

func (h *HeapImpl) getlen() int {
	return len(h.items)
}

func (h *HeapImpl) Peek() int {
	heaplen := h.getlen()
	if heaplen > 0{
		return h.items[0]
	}
	return 0
}
func main() {

	heap := NewHeap()

	heap.Push(2)
	heap.Push(3)
	heap.Push(1)

	heap.Pop()
	heap.Pop()

	fmt.Println(heap.items)
}