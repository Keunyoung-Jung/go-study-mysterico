package main

import (
	"fmt"
)

type List struct{
	head *Node
	tail *Node
}

type Node struct {
	prev *Node
	index int
	value interface{}
	next *Node
}

func (list *List) First()(*Node){
	return list.head
}

func (list *List) Append(value interface{}){
	node := &Node{value: value}
	if list.head == nil {
		list.head = node
		list.head.prev = nil
		list.head.index = 0
	} else {
		list.tail.next = node
		node.prev = list.tail
		node.index = list.tail.index +1
	}
	list.tail = node
}

func (node *Node) Next()(*Node){
	return node.next
}

func (node *Node) Prev()(*Node){
	return node.prev
}


func main() {
	list := List{}
	list.Append([]int{1,2})
	list.Append("hello")
	list.Append(1.5)

	mt.Println("{ prev   index   value   next       }")
	firstNode := list.First()
	fmt.Println(*firstNode)
	nextNode := firstNode.Next()
	fmt.Println(*nextNode)
	nextNode2 := nextNode.Next()
	fmt.Println(*nextNode2)

	fmt.Println("세번째 노드에서 이전으로 돌아가기")
	prevNode := nextNode2.Prev()
	fmt.Println(*prevNode)
}
