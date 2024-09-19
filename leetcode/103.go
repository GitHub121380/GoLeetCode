package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	left := true
	queue := make([]*TreeNode, 0, 2000)
	queue2 := make([]*TreeNode, 0, 2000)
	queue = append(queue, root)
	result := make([][]int, 0, 2000)
	var a []int
	for len(queue) != 0 {
		a = make([]int, 0, 1000)
		for len(queue) != 0 {
			a = append(a, queue[len(queue)-1].Val)
			if left {
				if queue[len(queue)-1].Left != nil {
					queue2 = append(queue2, queue[len(queue)-1].Left)
				}
				if queue[len(queue)-1].Right != nil {
					queue2 = append(queue2, queue[len(queue)-1].Right)
				}
			} else {
				if queue[len(queue)-1].Right != nil {
					queue2 = append(queue2, queue[len(queue)-1].Right)
				}
				if queue[len(queue)-1].Left != nil {
					queue2 = append(queue2, queue[len(queue)-1].Left)
				}
			}
			queue = queue[:len(queue)-1]
		}
		left = !left
		result = append(result, a)
		queue = queue2
		//queue2 = queue2[0:0]  //因为上面queue和queue2会指向同一个内存 如果这个时候queue2不新建 那么queue2 append的时候，就会影响queue
		queue2 = []*TreeNode{}
	}
	return result
}
func main() {

	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}
