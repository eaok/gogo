package main

import "fmt"

type Node struct {
	data int
	next *Node
}

//isCycle 判断是否有环
//使用快慢指针
func isCycle(head *Node) bool {
	p1, p2 := head, head

	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
		if p1 == p2 {
			return true
		}
	}

	return false
}

//cycleLength 返回环的长度
//快慢指针相遇后继续前进，统计到下一次相遇循环的次数
func cycleLength(head *Node) int {
	p1, p2 := head, head
	meetFlag := false
	count := 0

	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
		if p1 == p2 && meetFlag == false {
			meetFlag = true
		} else if p1 == p2 && meetFlag == true {
			break
		}

		if meetFlag == true {
			count++
		}
	}

	return count
}

//cycleNode 返回入环结点
//快慢指针相遇后，慢指针和新指针相遇后的结点就是
func cycleNode(head *Node) *Node {
	p1, p2, p3 := head, head, head
	meetFlag := false

	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
		if p1 == p2 {
			meetFlag = true
		}
		if meetFlag == true {
			if p1 == p3 {
				break
			}
			p3 = p3.next
		}
	}

	return p1
}

func main() {
	node1, node2, node3, node4, node5 := new(Node), new(Node), new(Node), new(Node), new(Node)
	node1.data = 5
	node2.data = 3
	node3.data = 7
	node4.data = 2
	node5.data = 6
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node2

	fmt.Println(isCycle(node1))
	fmt.Println(cycleLength(node1))
	fmt.Println(cycleNode(node1))
}