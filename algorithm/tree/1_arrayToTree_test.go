package tree

import (
	"fmt"
	"reflect"
	"testing"
)

// printTreeMidOrder 中序遍历树
func printTreeMidOrder(root *BNode) string {
	result := ""
	root.TraverseFunc(func(n *BNode) {
		result += fmt.Sprint(n.Data, " ")
	})
	fmt.Sprintln()

	return result
}

//TraverseFunc 中序遍历递归部分
func (node *BNode) TraverseFunc(f func(node *BNode)) {
	if node == nil {
		return
	}

	node.LeftChild.TraverseFunc(f)
	f(node)
	node.RightChild.TraverseFunc(f)
}

// func PrintTreeMidOrder(root *BNode) string {
// 	if root == nil {
// 		return
// 	}
// 	//遍历root结点的左子树
// 	if root.LeftChild != nil {
// 		PrintTreeMidOrder(root.LeftChild)
// 	}
// 	//遍历root结点
// 	fmt.Print(root.Data, " ")
// 	//遍历root结点的右子树
// 	if root.RightChild != nil {
// 		PrintTreeMidOrder(root.RightChild)
// 	}
// }

func Test_arrayToTree(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				arr:   []int{1, 2, 3, 4, 5},
				start: 0,
				end:   4,
			},
			want: "1 2 3 4 5 ",
		},
		{
			name: "case2",
			args: args{
				arr:   []int{1, 3, 5, 7, 9},
				start: 0,
				end:   4,
			},
			want: "1 3 5 7 9 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arrayToTree(tt.args.arr, tt.args.start, tt.args.end)
			println(tt.name, printTreeMidOrder(got))
			if !reflect.DeepEqual(printTreeMidOrder(got), tt.want) {
				t.Errorf("arrayToTree() = %v, want %v", printTreeMidOrder(got), tt.want)
			}
		})
	}
}
