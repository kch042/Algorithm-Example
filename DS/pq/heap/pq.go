package main

import (
	"errors"
	"fmt"
)

// a simple implementation of max priority queue
// using heap data structure

// ref: CLRS 6.5

type PQ []int

// need to use pointer to the slice
// since the append would return a new slice if need to expand
func (pq *PQ) Insert(key int) {
	*pq = append(*pq, key)
	pq.siftUp(len(*pq) - 1)
}

func (pq PQ) Increase(i, key int) error {
	if i < 0 || i >= len(pq) {
		return errors.New("Invalid index!")
	}
	if pq[i] > key {
		return errors.New("Input key less than the original!")
	}

	pq[i] = key
	pq.siftUp(i)
	return nil
}

func (pq PQ) siftUp(i int) {
	now := i
	for pa := (now - 1) / 2; pa >= 0; pa = (now - 1) / 2 {
		if pq[pa] >= pq[now] {
			return
		}
		pq[pa], pq[now] = pq[now], pq[pa]
		now = pa
	}
}

// since we resize the slice
// so it's better to use pointer to make sure the change
func (pq *PQ) Extract() (int, error) {
	if len(*pq) < 1 {
		return -1, errors.New("PQ underflow!")
	}

	ret := (*pq)[0]
	(*pq)[0], (*pq)[len(*pq)-1] = (*pq)[len(*pq)-1], (*pq)[0]
	*pq = (*pq)[:len(*pq)-1]
	pq.siftDown(0)

	return ret, nil
}

func (pq PQ) siftDown(root int) {
	for 2*root+1 < len(pq) {
		child := 2*root + 1
		if 2*root+2 < len(pq) && pq[2*root+2] > pq[child] {
			child = 2*root + 2
		}
		if pq[root] >= pq[child] {
			return
		}
		pq[root], pq[child] = pq[child], pq[root]
		root = child
	}
}

func main() {
	pq := make(PQ, 0)
	input := []int{5, 10, 85, 1, 3, 9, 4, 7, 100, 0}
	for _, i := range input {
		pq.Insert(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(pq.Extract())
	}
}
