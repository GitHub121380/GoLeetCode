package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 可以把left和right范围内的链表截断，当作一个
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if (left == 1 && right == 1) || head.Next == nil {
		return head
	}
	virtualHead := &ListNode{Val: 0, Next: head}

	//找到左、左的左、右、右的右
	var pre = virtualHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	leftp := pre.Next

	rightp := pre
	for i := 0; i < right-left+1; i++ {
		rightp = rightp.Next
	}
	curr := rightp.Next

	//切断
	pre.Next = nil
	rightp.Next = nil

	//逆转
	var before *ListNode
	head2 := leftp
	next := head2.Next
	for head2 != nil {
		head2.Next = before
		before = head2
		head2 = next
		if next != nil {
			//就算next.Next是nil 也没关系
			next = next.Next
		}
	}
	//退出来后，before就是被逆转的后的头结点，但我们不用管它，因为我们已经标记了left和right

	//连接
	pre.Next = rightp
	leftp.Next = curr

	return virtualHead.Next
}

// 在区间内采用头插法 在区域内从左到右，依次把所有结点插到left前面，直到执行次数达到right-left
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if (left == 1 && right == 1) || head.Next == nil {
		return head
	}

	virtualHead := &ListNode{Val: 0, Next: head}

	pre := virtualHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	next := cur.Next
	for i := 0; i < right-left; i++ {

		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
		next = cur.Next

	}
	return virtualHead.Next
}

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}
	fmt.Println(reverseBetween(head, 2, 4))
}
