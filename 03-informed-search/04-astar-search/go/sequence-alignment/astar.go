package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

// A simple priority queue for A* nodes
type PQItem struct {
	id    int     // node id = i*(cols) + j
	f     float64 // f = g + h
	g     float64 // cost so far
	index int
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].f < pq[j].f }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// astarAlign runs A* on the alignment grid and prints the resulting alignment
func astarAlign(ainstr, binstr string, mispen, gappen, skwpen float64) {
	ain := strings.Split(ainstr, "")
	bin := strings.Split(binstr, "")
	ia := len(ain)
	ib := len(bin)
	cols := ib + 1

	// helper to convert i,j to id
	id := func(i, j int) int { return i*cols + j }

	// admissible heuristic: assume all remaining paired positions are perfect matches (-1 each)
	heuristic := func(i, j int) float64 {
		ri := ia - i
		rj := ib - j
		pairs := ri
		if rj < ri {
			pairs = rj
		}
		// each match gives cost -1.0 (as in stringalign)
		return -1.0 * float64(pairs)
	}

	// cost to move
	moveCost := func(i, j, ni, nj int) float64 {
		// diag
		if ni == i+1 && nj == j+1 {
			if ain[i] == bin[j] {
				return -1.0
			}
			return mispen
		}
		// down (advance in A, gap in B)
		if ni == i+1 && nj == j {
			// if we're moving into last column (j == ib) apply skwpen
			if j == ib {
				return skwpen
			}
			return gappen
		}
		// right (gap in A, advance in B)
		if ni == i && nj == j+1 {
			if i == ia {
				return skwpen
			}
			return gappen
		}
		return math.Inf(1)
	}

	start := id(0, 0)
	goal := id(ia, ib)

	// open set pq and maps
	open := make(map[int]*PQItem)
	cameFrom := make(map[int]int)
	// move type: 0 diag, 1 down (A), 2 right (B)
	moveType := make(map[int]int)
	gScore := make(map[int]float64)

	pq := &PriorityQueue{}
	heap.Init(pq)

	startItem := &PQItem{id: start, g: 0, f: heuristic(0, 0)}
	heap.Push(pq, startItem)
	open[start] = startItem
	gScore[start] = 0

	expanded := 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PQItem)
		delete(open, current.id)
		expanded++

		if current.id == goal {
			// reconstruct path
			reconstructAndPrint(current.id, cameFrom, moveType, ain, bin, ia, ib, current.g)
			fmt.Printf("Nodes expanded: %d\n", expanded)
			return
		}

		// decode i,j
		ci := current.id / cols
		cj := current.id % cols

		// neighbors: diag, down, right
		neighbors := [][2]int{{ci + 1, cj + 1}, {ci + 1, cj}, {ci, cj + 1}}
		for k, nb := range neighbors {
			ni, nj := nb[0], nb[1]
			if ni < 0 || nj < 0 || ni > ia || nj > ib {
				continue
			}
			nid := id(ni, nj)
			cost := moveCost(ci, cj, ni, nj)
			tentativeG := current.g + cost

			oldG, ok := gScore[nid]
			if !ok || tentativeG < oldG {
				cameFrom[nid] = current.id
				moveType[nid] = k // 0 diag,1 down,2 right
				gScore[nid] = tentativeG
				f := tentativeG + heuristic(ni, nj)
				if item, exists := open[nid]; exists {
					item.g = tentativeG
					item.f = f
					heap.Fix(pq, item.index)
				} else {
					item := &PQItem{id: nid, g: tentativeG, f: f}
					heap.Push(pq, item)
					open[nid] = item
				}
			}
		}
	}

	fmt.Println("No se encontró camino (imposible)")
}

func reconstructAndPrint(goal int, cameFrom map[int]int, moveType map[int]int, ain, bin []string, ia, ib int, totalCost float64) {
	cols := ib + 1
	// build path
	path := []int{goal}
	cur := goal
	for cur != 0 {
		prev, ok := cameFrom[cur]
		if !ok {
			break
		}
		path = append(path, prev)
		cur = prev
	}
	// reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	var aout []string
	var bout []string

	for idx := 1; idx < len(path); idx++ {
		prev := path[idx-1]
		cur := path[idx]
		pi := prev / cols
		pj := prev % cols
		ci := cur / cols
		cj := cur % cols
		if ci == pi+1 && cj == pj+1 {
			aout = append(aout, ain[pi])
			bout = append(bout, bin[pj])
		} else if ci == pi+1 && cj == pj {
			aout = append(aout, ain[pi])
			bout = append(bout, " ")
		} else if ci == pi && cj == pj+1 {
			aout = append(aout, " ")
			bout = append(bout, bin[pj])
		}
	}

	fmt.Println("\nA* Alineamiento A:")
	fmt.Println(strings.Join(aout, ""))
	fmt.Println("A* Alineamiento B:")
	fmt.Println(strings.Join(bout, ""))
	fmt.Printf("A* Coste total: %.2f\n", totalCost)
}
