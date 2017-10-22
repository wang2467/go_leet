package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
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

func topologicalSort_top(g [][]int) *arraystack.Stack {
	visited := make([]int, len(g))
	for i := range visited {
		visited[i] = 0
	}
	result := arraystack.New()
	for i := range g {
		if visited[i] == 0 {
			topologicalSort(g, visited, result, i)
		}
	}
	for !result.Empty() {
		x, _ := result.Pop()
		fmt.Println(x)
	}
	return result
}

func topologicalSort(g [][]int, visited []int, result *arraystack.Stack, idx int) {
	visited[idx] = 1
	for i := range g[idx] {
		if visited[g[idx][i]] == 0 {
			topologicalSort(g, visited, result, g[idx][i])
		}
	}
	result.Push(idx)
}

func main() {
	g := makeGraph(6)
	addEdges(g, 5, 2)
	addEdges(g, 5, 0)
	addEdges(g, 4, 1)
	addEdges(g, 4, 0)
	addEdges(g, 2, 3)
	addEdges(g, 3, 1)
	topologicalSort_top(g)
}
