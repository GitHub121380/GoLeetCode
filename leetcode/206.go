package main

/*
*
  - Definition for singly-linked list.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var behind *ListNode = nil
	next := head.Next

	for head != nil {
		head.Next = behind
		behind = head
		head = next
		if next != nil {
			next = next.Next
		}
	}
	return behind
}

// 递归法 本质其实是用栈先入后出的特点来做
// 既然先入后出 那么一定是尾部的先被逆转 故函数整体代表的是从head往后的已经被逆的尾部，head指向了逆转后的尾（最靠近head的那个），而要弹出前我们针对栈顶结点做指针逆转操作，站在栈顶的视角想办法把已经逆转后的链接上自己
// head.next正序的下一个结点，那么让这个结点去指向head，就成为了head的父节点。目前head指向了父节点，父节点也指向head，循环了，把head.next置空，就断开了循环
// 但是因为输出是要逆转后的头结点，所以我们只能一直return逆转的头；而不是return逆转的尾；压到最后
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
