package main

import "fmt"

type node struct {
	// id   int
	visited bool
	pi      *node
	d, f    int // discovered time, finish time
}

// unweighted directed graph
type graph struct {
	edges map[int][]int
	nodes map[int]*node
}

func newGraph() *graph {
	g := new(graph)
	g.edges, g.nodes = make(map[int][]int), make(map[int]*node)
	return g
}

func (g *graph) addEdge(u, v int) {
	if _, ok := g.nodes[u]; !ok {
		g.nodes[u] = &node{visited: false}
	}
	if _, ok := g.nodes[v]; !ok {
		g.nodes[v] = &node{visited: false}
	}
	g.edges[u] = append(g.edges[u], v)
}

func (g *graph) reset() {
	for _, node := range g.nodes {
		node.visited = false
		node.pi = nil
	}
}

func (g *graph) dfs(id, time int) int {
	g.nodes[id].d = time
	g.nodes[id].visited = true

	for _, adj := range g.edges[id] {
		if !g.nodes[adj].visited {
			g.nodes[adj].pi = g.nodes[id]
			time = g.dfs(adj, time+1)
		}
	}
	time++
	g.nodes[id].f = time
	return time
}

func (g *graph) DFS(order ...int) {
	if len(order) < len(g.nodes) { // didn't specify the order
		order = make([]int, 0, len(g.nodes))
		for k := range g.nodes {
			order = append(order, k)
		}
	}

	time := 1
	for _, next := range order {
		if !g.nodes[next].visited {
			time = g.dfs(next, time) + 1
		}
	}
}

func partition(lo, hi int, nodes []*node) int {
	x, i := nodes[lo], lo
	for j := lo + 1; j < hi; j++ {
		if x.f > nodes[j].f {
			i++
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	nodes[lo], nodes[i] = nodes[i], nodes[lo]
	return i
}

func sortFinish(lo, hi int, nodes []*node) {
	if lo < hi {
		m := partition(lo, hi, nodes)
		sortFinish(lo, m, nodes)
		sortFinish(m+1, hi, nodes)
	}
}

func (g *graph) getDecreasingFinish() []int {
	nodes := make([]*node, 0, len(g.nodes))
	for _, node := range g.nodes {
		nodes = append(nodes, node)
	}

	sortFinish(0, len(nodes), nodes)

	ret := make([]int, len(nodes))
	for i, v := range nodes {
		ret[i] = v.f // problem
	}

	return nil
}

func (g *graph) transpose() *graph {
	ret := &graph{nodes: g.nodes, edges: make(map[int][]int)}
	for node := range g.nodes {
		for _, v := range g.edges[node] {
			ret.edges[v] = append(ret.edges[v], node)
		}
	}

	return ret
}

func (g *graph) sccDFS(order []int) [][]int {
	ret := make([][]int, 0)
	g.reset()
	i := 0

	// do simple dfs
	var dfs func(id int)
	dfs = func(id int) {
		g.nodes[id].visited = true
		ret[i] = append(ret[i], id)
		for _, adj := range g.edges[id] {
			if !g.nodes[adj].visited {
				dfs(adj)
			}
		}
	}

	for _, v := range order {
		if !g.nodes[v].visited {
			dfs(v)
			i++
		}
	}

	return ret
}

func (g *graph) SCC() [][]int {
	g.DFS(0)

	order := g.getDecreasingFinish()
	gt := g.transpose()

	return gt.sccDFS(order)
}

func (g *graph) Seed() {
	g.addEdge(0, 1)
	g.addEdge(1, 2)
	g.addEdge(1, 4)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(2, 5)
	g.addEdge(3, 2)
	g.addEdge(4, 5)
	g.addEdge(4, 6)
	g.addEdge(5, 4)
	g.addEdge(5, 6)
	g.addEdge(5, 7)
	g.addEdge(6, 7)
	g.addEdge(7, 8)
	g.addEdge(8, 6)
}

func main() {
	g := newGraph()
	g.Seed()
	groups := g.SCC()
	for i, c := range groups {
		fmt.Printf("#%d: ", i)
		for _, mem := range c { // group member
			fmt.Printf("%d ", mem)
		}
		fmt.Println()
	}
}
