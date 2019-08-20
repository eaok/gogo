package list

//LNode 结点
type LNode struct {
	Data interface{}
	Next *LNode
}

//ReverseV1 就地逆序
//遍历链表时，让当前结点指向前驱结点
func ReverseV1(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *LNode
	var cur *LNode
	next := node.Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}

	node.Next = pre
}

//ReverseV2 递归实现逆序
func ReverseV2(node *LNode) {
	firstNode := node.Next
	newHead := RecursiveReverse(firstNode)
	node.Next = newHead
}

//RecursiveReverse 递归的函数
func RecursiveReverse(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	newHead := RecursiveReverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

//ReverseV3 插入法
//固定把第二个结点插入到头结点后面
func ReverseV3(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	node.Next.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}
