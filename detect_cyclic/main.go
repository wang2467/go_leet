package main

import (
	"fmt"
)

func addEdges(g [][]int, a int, b int) {
	g[a] = append(g[a], b)
}

func makeGraph(a int) [][]int {
	g := make([][]int, a)
	for i := range g {
		g[i] = make([]int, 0)
	}
	return g
}

func detectCyclics(g [][]int, visited []int, recur []int, idx int) bool {
	visited[idx] = 1
	recur[idx] = 1
	for i := range g[idx] {
		if visited[g[idx][i]] == 0 {
			if detectCyclics(g, visited, recur, g[idx][i]) == true {
				return true
			}
		} else if recur[g[idx][i]] == 1 {
			return true
		}
	}
	recur[idx] = 0
	return false
}

func detectCyclics_top(g [][]int) bool {
	visited := make([]int, len(g))
	recur := make([]int, len(g))
	for i := range g {
		if visited[i] == 0 {
			if detectCyclics(g, visited, recur, i) {
				return true
			}
		}
	}
	return false
}

func main() {
	g := makeGraph(4)
	addEdges(g, 0, 1)
	addEdges(g, 0, 2)
	addEdges(g, 1, 2)
	addEdges(g, 2, 0)
	addEdges(g, 2, 3)
	addEdges(g, 3, 3)
	fmt.Println(detectCyclics_top(g))
}
