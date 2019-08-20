package list

import (
	"fmt"
	"testing"
)

//CreateNode 创建一个链表
func CreateNode(node *LNode, num int) {
	cur := node
	for i := 0; i < num; i++ {
		cur.Next = &LNode{Data: i}
		cur = cur.Next
	}
}

//PrinList 显示链表中所有元素
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func TestReverseV1(t *testing.T) {
	head1 := &LNode{}
	head2 := &LNode{}
	CreateNode(head1, 3)
	CreateNode(head2, 7)
	type args struct {
		node *LNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "就地逆序head1",
			args: args{
				head1,
			},
		},
		{
			name: "就地逆序head2",
			args: args{
				head2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNode(tt.name, tt.args.node)
			ReverseV1(tt.args.node)
			PrintNode(tt.name, tt.args.node)
		})
	}
}

func TestReverseV2(t *testing.T) {
	head1 := &LNode{}
	head2 := &LNode{}
	CreateNode(head1, 3)
	CreateNode(head2, 7)
	type args struct {
		node *LNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "递归法head1",
			args: args{
				head1,
			},
		},
		{
			name: "递归法head2",
			args: args{
				head2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNode(tt.name, tt.args.node)
			ReverseV2(tt.args.node)
			PrintNode(tt.name, tt.args.node)
		})
	}
}

func TestReverseV3(t *testing.T) {
	head1 := &LNode{}
	head2 := &LNode{}
	CreateNode(head1, 3)
	CreateNode(head2, 7)
	type args struct {
		node *LNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "插入法head1",
			args: args{
				head1,
			},
		},
		{
			name: "插入法head2",
			args: args{
				head2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintNode(tt.name, tt.args.node)
			ReverseV3(tt.args.node)
			PrintNode(tt.name, tt.args.node)
		})
	}
}
