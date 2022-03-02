package main

import (
	"fmt"
)

const inf int = (1 << 31) - 1

// this only applies to connected graph
type graph struct {
	w map[int]map[int]int
}

func newGraph() *graph {
	g := new(graph)
	(*g).w = make(map[int]map[int]int)
	return g
}

func (g *graph) AddEdge(u, v, w int) {
	if len(g.w[u]) == 0 {
		g.w[u] = make(map[int]int)
	}
	if len(g.w[v]) == 0 {
		g.w[v] = make(map[int]int)
	}
	g.w[u][v] = w
	g.w[v][u] = w
}

// -----------------------------------------------------

func (g *graph) dfs(s int) {
	if _, ok := g.w[s]; !ok {
		fmt.Println("Starting node is not in the graph!")
		return
	}

	visited := make(map[int]bool)
	for k := range g.w {
		visited[k] = false
	}

	time := 0
	var search func(s int)
	search = func(s int) {
		visited[s] = true
		time++
		fmt.Printf("Discover %d at time %d\n", s, time)
		for nei := range g.w[s] {
			if !visited[nei] {
				search(nei)
			}
		}
		time++
		fmt.Printf("Finish %d at time %d\n", s, time)
	}
	search(s)
	fmt.Println()
}

// for dfs
func (g *graph) Seed() {
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1)
	g.AddEdge(1, 5, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 2, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(7, 6, 1)
	g.AddEdge(5, 8, 1)
	g.AddEdge(9, 8, 1)
	g.AddEdge(9, 10, 1)
	g.AddEdge(10, 8, 1)
	g.AddEdge(8, 1, 1)
}

// -------------------------------------------------

type node struct {
	id  int   // the id of the node
	val int   // for constructing MST
	pi  *node // predecessor
}

type heap struct {
	data [](*node)
	find map[int]int // find[k] returns the index of the node with key k in the heap
}

func (h *heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.find[h.data[i].id] = i
	h.find[h.data[j].id] = j
}

func (h *heap) siftUp(now int) {
	if now <= 0 || now >= len(h.data) {
		return
	}
	for (now-1)/2 >= 0 && now > 0 {
		p := (now - 1) / 2
		if h.data[now].val >= h.data[p].val {
			return
		}
		h.swap(now, p)
		now = p
	}
}

func (h *heap) extractMin() *node {
	ret := h.data[0]
	h.swap(0, len(h.data)-1)
	delete(h.find, ret.id) // need to delete after insertion

	h.data = h.data[:len(h.data)-1]
	h.siftDown(0)

	return ret
}

func (h *heap) siftDown(now int) {
	if now < 0 || now >= len(h.data) {
		return
	}
	for 2*now+1 < len(h.data) {
		child := 2*now + 1
		if child+1 < len(h.data) && h.data[child+1].val < h.data[child].val {
			child++
		}
		if h.data[child].val >= h.data[now].val {
			return
		}
		h.swap(now, child)
		now = child
	}
}

func (g *graph) newHeap() *heap {
	h := &heap{}
	h.data = make([]*node, len(g.w))
	h.find = make(map[int]int)

	i := 0
	for id := range g.w {
		h.data[i] = &node{id: id, val: inf}
		h.find[id] = i
		i++
	}

	for j := (len(g.w) - 2) / 2; j >= 0; j-- {
		h.siftDown(j)
	}
	return h
}

func (h *heap) decreaseVal(neiID, newVal int, pi *node) {
	if h.isInHeap(neiID) {
		at := h.find[neiID]
		if h.data[at].val > newVal {
			h.data[at].val = newVal
			h.data[at].pi = pi
			h.siftUp(at)
		}
	}
}

func (h *heap) isInHeap(id int) bool {
	_, ok := h.find[id]
	return ok
}

func (h *heap) getNode(id int) *node {
	return h.data[h.find[id]]
}

// Prim's algorithm
func (g *graph) MST(stID int) [](*node) {
	// make new heap and decrease key for starting node
	h := g.newHeap()
	h.decreaseVal(stID, 0, nil)
	ret := make([](*node), 0, len(g.w))

	for len(h.data) > 0 {
		now := h.extractMin()
		ret = append(ret, now)
		for neiID, wgt := range g.w[now.id] {
			h.decreaseVal(neiID, wgt, now)
		}
	}

	return ret
}

func printMST(nodes [](*node)) {
	for _, n := range nodes {
		for n.pi != nil {
			fmt.Printf("[%d]---%d--->", n.id, n.val)
			n = n.pi
		}
		fmt.Println(n.id)
	}
}

// for MST
func (g *graph) Seed2() {
	g.AddEdge(1, 2, 15)
	g.AddEdge(1, 3, 10)
	g.AddEdge(1, 4, 7)
	g.AddEdge(3, 5, 3)
	g.AddEdge(3, 6, 8)
	g.AddEdge(5, 6, 14)
	g.AddEdge(4, 6, 5)
	g.AddEdge(4, 7, 12)
	g.AddEdge(4, 8, 9)
	g.AddEdge(6, 7, 6)
}

func main() {
	g := newGraph()
	g.Seed2()
	printMST(g.MST(1))

}
