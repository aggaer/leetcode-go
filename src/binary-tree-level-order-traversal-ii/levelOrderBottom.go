package main

import "fmt"

var treeRoot TreeNode

func main() {
	if &treeRoot == nil {
		panic("tree root is nil")
	}
	//traversal(&treeRoot)
	orders := levelOrderBottom(&treeRoot)
	for _, v := range orders {
		for _, v := range v {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var tempRet [][]int
	if root == nil {
		return tempRet
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		var stackTemp []*TreeNode
		var tempRetList []int
		for _, v := range stack {
			tempRetList = append(tempRetList, v.Val)
			if v.Left != nil {
				stackTemp = append(stackTemp, v.Left)
			}
			if v.Right != nil {
				stackTemp = append(stackTemp, v.Right)
			}
		}
		tempRet = append(tempRet, tempRetList)
		stack = stackTemp
	}
	ret := make([][]int, len(tempRet))
	for i := range tempRet {
		ret[i] = tempRet[len(tempRet)-i-1]
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(root *TreeNode) {
	if root != nil {
		fmt.Printf(" %d", root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
}

func init() {
	treeRoot = TreeNode{Val: 3}
	treeRoot.Left = &TreeNode{Val: 9}
	treeRoot.Right = &TreeNode{Val: 20}
	treeRoot.Right.Left = &TreeNode{Val: 15}
	treeRoot.Right.Right = &TreeNode{Val: 7}
}
