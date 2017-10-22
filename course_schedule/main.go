package main

import (
	"fmt"
)

func findOrder(numCourse int, prerequisite [][]int) []int {
	g := make([][]int, numCourse)
	for i := range g {
		g[i] = make([]int, 0)
	}
	for i := range prerequisite {
		p := prerequisite[i]
		g[p[0]] = append(g[p[0]], p[1])
	}
	return topologicalSort_top(g)
}

func topologicalSort_top(g [][]int) []int {
	visited := make([]int, len(g))
	result := make([]int, 0)
	recur := make([]int, len(g))
	for i := range g {
		if visited[i] == 0 {
			topologicalSort(g, visited, recur, i, &result)
		}
	}
	return result
}

func topologicalSort(g [][]int, visited []int, recur []int, idx int, result *[]int) bool {
	visited[idx] = 1
	recur[idx] = 1
	for i := range g[idx] {
		if visited[g[idx][i]] == 0 {
			if topologicalSort(g, visited, recur, g[idx][i], result) {
				*result = nil
				return true
			}
		} else if recur[g[idx][i]] == 1 {
			*result = nil
			return true
		}
	}
	recur[idx] = 0
	*result = append(*result, idx)
	return false
}

func main() {
	s := make([][]int, 4)
	s[0] = []int{1, 0}
	s[1] = []int{2, 0}
	s[2] = []int{3, 1}
	s[3] = []int{3, 2}
	fmt.Println(findOrder(4, s))
}
