package tree

import (
	"fmt"
	"strconv"
	"strings"
)

//TreeNode is a struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree return a *TreeNode
func BuildTree() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Left = &TreeNode{
		Val: 4,
	}
	root.Right.Right = &TreeNode{
		Val: 5,
	}
	return root
}

// PrintTree 层序打印一棵树
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println(root)
		return
	}
	queue := []*TreeNode{root}
	res := []string{}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "#")
			continue
		}
		res = append(res, strconv.Itoa(node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	fmt.Println(res)
}

// Serialize 二叉树序列化-递归
func Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + Serialize(root.Left) + "," + Serialize(root.Right)
}

// UnSerialize 二叉树反序列化-递归
func UnSerialize(str string) *TreeNode {
	fields := strings.Split(str, ",")
	// 需要一个变量每执行一次就增加一次，需要是全局变量，故不能放到递归参数中
	index := 0
	var dfs func([]string) *TreeNode
	dfs = func(s []string) *TreeNode {
		if index > len(s)-1 {
			return nil
		}
		value := s[index]
		index++
		if value != "#" {
			valueInt, _ := strconv.Atoi(value)
			root := &TreeNode{
				Val: valueInt,
			}
			root.Left = dfs(s)
			root.Right = dfs(s)
			return root
		}
		return nil
	}
	tree := dfs(fields)
	return tree
}

// Serialize1 二叉树序列化-层序
func Serialize1(root *TreeNode) string {
	res := []string{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, "#")
			continue
		}
		res = append(res, strconv.Itoa(node.Val))
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return strings.Join(res, ",")
}
