package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 逐层遍历 bfs 队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0, 2000)
	l1 := list.New()
	l1.PushBack(root)
	l2 := list.New()
	//提前声明，避免重复申请内存，leetcode上耗时可从3ms降低到0ms
	var node *TreeNode
	for l1.Len() != 0 {
		s := make([]int, 0, 32)
		for l1.Len() != 0 {
			node = l1.Front().Value.(*TreeNode)
			if node.Left != nil {
				l2.PushBack(node.Left)
			}
			if node.Right != nil {
				l2.PushBack(node.Right)
			}
			s = append(s, node.Val)
			l1.Remove(l1.Front())
		}

		l1.Init()
		l1.PushBackList(l2)
		l2.Init()
		result = append(result, s)
	}
	return result
}

// dfs 深度优先 但是逐层输出，相当于每层第一个结点去append一个[]int并往里append val
// 之后这层的结点执行递归函数时，找到那个[]int，去往里append val
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0, 2000)

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		//len表示已经走过了几层，depth表示当前是第几层，另外result的每个元素下标就是对应的层 root是0层
		//因为depth从0开始，所以真正到的层数其实是depth+1，比如root的孙子结点，depth=2，此时len=2
		if len(result) < depth+1 {
			result = append(result, make([]int, 0, 32))
		}
		result[depth] = append(result[depth], root.Val)
		//去下一层
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return result
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}

	fmt.Println(levelOrder(root))

}
