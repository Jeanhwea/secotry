package ch02_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(vals []int) *ListNode {
	hair := &ListNode{}
	tail := hair

	for _, val := range vals {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}

	return hair.Next
}

func ListToArr(head *ListNode) []int{
	var arr []int
	for p := head; p != nil ; p = p.Next {
		arr = append(arr, p.Val)
	}
	return arr
}
