package main

// 如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
// 因此，该问题可以转化为：两个树在什么情况下互为镜像？
// 如果同时满足下面的条件，两个树互为镜像：
//
//	它们的两个根结点具有相同的值
//	每个树的右子树都与另一个树的左子树镜像对称
//
// 我们可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，p 指针和 q 指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，q 右移。每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称。
// 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	//两个都到了叶子的下一层
	if left == nil && right == nil {
		return true
	}

	//一个为空 一个不为空
	if left == nil || right == nil {
		return false
	}

	//自己相同了  还要左的左和右的右相同 还要左的右和右的左相同
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 递归
// 引入队列
// 初始化的时候把根节点入队两次，然后每次出两个，比较值，然后把第一个节点从左开始入队，第二个结点从右开始入队，一次一个
// 总之就是左边的从左向右 右边的 从右向左
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	list := make([]*TreeNode, 0, 1000)
	list = append(list, root)
	list = append(list, root)
	for len(list) != 0 {
		//p q是镜像的一左一右
		p, q := list[0], list[1]
		//等价于出队
		list = list[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}

		//	入队也保持对称
		//	p入左q入右
		list = append(list, p.Left, q.Right)
		//p入右q入左
		list = append(list, p.Right, q.Left)

		//list = append(list, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}
