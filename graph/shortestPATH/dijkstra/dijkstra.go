// Key Idea of Dijkstra's Algorithm:
// If u is the next node to be picked up by the algorithm
// then dist[u] must be the shortest path length
// from s to u at the moment we pick u up

package main

import (
	"fmt"
	"math"
)

// pseudocode
// G: Graph, s: starting vertex, R: visited vertexs, V: All vertexs

// Dijkstra(G, s):
//   for each u in V\{s}:
//   	d[u] = infinity
//   d[s] = 0
//   R = []
//   for V != R:
//     select u not in R s.t. d[u] is smallest
//     R = R U {u}
//     for each neighbor v of u:
//       if d[v] > d[u] + len(u, v):
//   	 	d[v] = d[u] + len(u, v)

// Dijkstra(G, s):
//   for each u in G:
//   	dist[u] = inf
//   dist[s] = 0
//
//   q := {s}
//   for q is not empty:
//     pop = q[0]
//     q = q[1:]
//     for each neighbor nei of pop:
//       if dist[nei] > dist[pop] + len(pop, nei):
//		   dist[nei] = dist[pop] + len(pop, nei)
//		   add nei to q
//    return dist[end]

// Using Priority Queue (can be implemented using heap packge)
// Dijkstra(G, s):
//   pq := {node{v: s, dist: 0}}
//   visited := map[int]struct{}
//   for pq not empty:
//     pop := pq.Pop()
//	   if pop == end:
//       return pop.dist
//     for each neighbor nei of pop:
//       if nei is not visited:
//       	if nei.dist > pop.dist + len(pop, nei):
//				add {node{nei, pop.dist+len(pop, nei)}} to pq
//				heapify
//	  return -1  (Not Found)

// Note that using pq, once the end node is picked up
// it is guaranteed to have smallest dist from src to end since pq is a minHeap

func main() {
	times := make([][]int, 1)
	times[0] = []int{1, 2, 1}
	fmt.Println(networkDelayTime(times, 2, 2))
}

type dis struct {
	dest int
	d    float64
}

func networkDelayTime(times [][]int, N int, K int) int {
	graph := make(map[int][]dis)
	for _, t := range times {
		add := dis{dest: t[1], d: float64(t[2])}
		graph[t[0]] = append(graph[t[0]], add)
	}

	visited := make(map[int]struct{})
	dist := make([]float64, N+1)
	for i := range dist {
		dist[i] = math.Inf(1)
	}

	dist[K] = 0
	for {
		now := -1
		now_dist := math.Inf(1)

		// Extract min
		for i := 1; i <= N; i++ {
			if _, ex := visited[i]; !ex {
				if now_dist > dist[i] {
					now_dist = dist[i]
					now = i
				}
			}
		}

		if now == -1 {
			break
		}

		visited[now] = struct{}{}
		for _, nei := range graph[now] {
			if dist[nei.dest] > now_dist+nei.d {
				dist[nei.dest] = now_dist + nei.d
			}
		}
	}

	return max(dist)
}

func max(dist []float64) int {
	ret := math.Inf(-1)
	for i := 1; i < len(dist); i++ {
		if dist[i] == math.Inf(1) {
			return -1
		}
		if dist[i] > ret {
			ret = dist[i]
		}
	}

	return int(ret)
}

// Correctness of Dijkstra's Algorithm

// Prove by induction
// if only one vertex u, which is trivial
// -> d[u] = 0, Correct

// Denote 𝛿[u] as the the length of the shortest path from s to u
//        d[u] be the distance calculated by dijkstra's algorithm

// Induction Hypothesis (I.H.)
//   Let R be the collection of visited vertexs with 𝛿[u] = d[u]
//   Let u' be the last vertex added to R, we need to show that
//   				𝛿[u'] = d[u']

// Prove by contradiction
//   Suppose there's a path Q from s to u' with
//   		len(Q) < d[u']

// Let x-y be the first edge that leaves R
// (i.e. the first edge that does not consists of vertex in R)
// then we have
// 		len(Q_x) + len(x, y) <= len(Q) < d[u'],
// where Q_x is the subpath of Q from s to x

// since x is in R, by I.H, we have
//      d[x] <= len(Q_x)
// so d[x] + len(x, y) < d[u']
// Also since y is adjacent to x
// d[y] is relaxed (if necessary) by the algorithm when visiting x s.t.
// 		d[y] <= d[x] + len(x, y)
// Hence we have
// 		d[y] < d[u']  --------- A

// Howerver, since the algorithm always picks the vertex with smallest d
// and u is picked by the algorithm while v isn't
// so u must have smallest distance label, that is
// 		d[u] <= d[y]  --------- B

// By A and B, we obtain an contradiction
// hence no such path as Q exists
// so 𝛿[u'] = d[u']
