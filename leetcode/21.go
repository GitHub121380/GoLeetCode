package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 新头结点+双指针遍历比较
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	newHead := &ListNode{Val: 0, Next: nil}
	current := newHead
	i := list1
	j := list2

	for i != nil && j != nil {
		if i.Val < j.Val {
			current.Next = i
			i = i.Next
		} else {
			current.Next = j
			j = j.Next
		}
		current = current.Next
	}
	if i == nil {
		current.Next = j
	} else if j == nil {
		current.Next = i
	}

	return newHead.Next
}

// 递归  函数本身就表示一个已经排好序的结果链表 发现自己元素小， 就把自己提出来，然后让自己的子结点链表跟另一个做排序
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
