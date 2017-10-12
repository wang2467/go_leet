package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func findMax(nums []int, x int, y int) int {
	max := nums[x]
	var idx int = x
	for i := x; i <= y; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}
	return idx
}

func buildMaxTree(nums []int) *tNode {
	return _buildMaxTree(nil, nums, 0, len(nums)-1)
}

func _buildMaxTree(root *tNode, nums []int, x int, y int) *tNode {
	if x > y {
		return nil
	}
	max_idx := findMax(nums, x, y)
	root = &tNode{val: nums[max_idx], left: nil, right: nil}
	root.left = _buildMaxTree(root, nums, x, max_idx-1)
	root.right = _buildMaxTree(root, nums, max_idx+1, y)
	return root
}

func printTree(root *tNode) {
	if root == nil {
		return
	}
	printTree(root.left)
	fmt.Println(root.val)
	printTree(root.right)
}
func main() {
	//fmt.Println(findMax([]int{1, 2, 5, 3}))
	r := buildMaxTree([]int{3, 2, 1, 6, 0, 5})
	printTree(r)
	//fmt.Println(([]int{1, 2, 3, 4})[0:3])
}
