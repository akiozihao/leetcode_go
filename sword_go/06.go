package swordgo

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	n, p := 0, head
	for p != nil {
		n++
		p = p.Next
	}
	res := make([]int, n)
	cur, p := n-1, head
	for p != nil {
		res[cur] = p.Val
		p = p.Next
		cur--
	}
	return res
}
