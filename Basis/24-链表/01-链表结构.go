package main

import "fmt"

type SingleNode struct {
	Data int
	Next *SingleNode
}

type SingleNodeList struct {
	First *SingleNode
	Last  *SingleNode
	Size  int
}

func NewSingleNodeList() *SingleNodeList {
	return &SingleNodeList{}
}

func (list *SingleNodeList) AddOneNodeAtTheLast(value int) {
	newNode := new(SingleNode)
	newNode.Data = value

	if list.Size == 0 {
		list.First = newNode
		list.Last = newNode
	} else {
		// 原最后一个节点->新节点
		lastNode := list.Last
		lastNode.Next = newNode
		newNode.Next = nil

		// 原链表的last->新节点
		list.Last = newNode
	}
	list.Size += 1
}

// 最后一个节点 = nil
// 原链表的last = nil，倒数第二个node->nil
func (list *SingleNodeList) DeleteOneNodeAtTheLast() {
	if list.Size == 0 {
		return
	} else if list.Size == 1 {
		list.First = nil
		list.Last = nil
		list.Size = 0
		return
	}

	var lastNode *SingleNode
	// node1->node2->node3
	currentNode := list.First
	for i := 0; i < list.Size; i++ {
		if i > 0 {
			currentNode = currentNode.Next
			if i == list.Size-2 {
				lastNode = currentNode
			}
		}
	}

	lastNode.Next = nil
	list.Last = lastNode
	list.Size -= 1
}

func (list *SingleNodeList) PrintNodeList() {
	currentNode := list.First
	fmt.Print(currentNode.Data, ",")
	for i := 0; i < list.Size; i++ {
		if currentNode.Next == nil {
			return
		}
		currentNode = currentNode.Next
		fmt.Print(currentNode.Data, ",")
	}
}

func main() {
	nodeList := NewSingleNodeList()
	nodeList.AddOneNodeAtTheLast(1)
	nodeList.AddOneNodeAtTheLast(4)
	nodeList.AddOneNodeAtTheLast(2)
	nodeList.PrintNodeList()
	nodeList.AddOneNodeAtTheLast(8)
	nodeList.PrintNodeList()

}
