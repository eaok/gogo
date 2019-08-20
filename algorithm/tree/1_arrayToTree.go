package tree

//BNode 二叉树结点
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

//NewBNode 新建一个二叉树
func NewBNode() *BNode {
	return &BNode{}
}

//arrayToTree 把一个有序整数数组放到二叉树中
func arrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = NewBNode()
		mid := (start + end + 1) / 2
		root.Data = arr[mid]
		root.LeftChild = arrayToTree(arr, start, mid-1)
		root.RightChild = arrayToTree(arr, mid+1, end)
	}

	return root
}
