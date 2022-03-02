package main

import (
	"container/heap"
	"fmt"
)

// Note when implementing a interface
// the method using value can be considered for both pointer and value
// while the one using pointer can only be for pointer

// For example
// type foo int
// type fooler interface {fool()}
// func (f foo) fool() {}
// ----> both foo and *foo fulfill the interface

// however if the fool() changes to
// func (f *foo) fool()
// ----> only *foo implement the interface

// Item is
type Item struct {
	name string
	ind  int
	pr   int
}

type PQ []Item

// Less decides the minHeap or maxHeap
// below is the minHeap case (maxHeap just reverse the inequality)
func (pq PQ) Less(i, j int) bool {
	return pq[i].pr < pq[j].pr
}

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].ind = i
	pq[j].ind = j
}

// we implement Pop and Push method
// using pointer to modify the pq correctly
// thus we have implement the interface in *pq not pq
func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	item.ind = -1
	*pq = (*pq)[:n-1]
	return item
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(Item) // type assertion for interface
	item.ind = n + 1
	*pq = append(*pq, item)
}

func main() {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"orange": 5,
	}

	i := 0
	pq := make(PQ, 3)
	for name, pr := range items {
		pq[i] = Item{name, i, pr}
		i++
	}
	heap.Init(&pq)

	newItem := Item{name: "pear", pr: 6}
	heap.Push(&pq, newItem)

	for pq.Len() > 0 {
		pop := heap.Pop(&pq).(Item)
		fmt.Printf("name: %s, pr: %d\n", pop.name, pop.pr) // name: orange, pr: 6
	}
}
