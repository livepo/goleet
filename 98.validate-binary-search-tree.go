package main

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}



func isValidBST(root *TreeNode) bool {
	if root == nil { return True }
	
	leftMax < root.Val < rightMin

}
