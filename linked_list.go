package linked_list

import "github.com/tjkomor/linked-list/node"

type LinkedList struct {
	Head node.Node
	Tail node.Node
}

func (self *LinkedList) Append(data string) {
	newNode := node.Node{Data: data}
	// println(new_node.Data)
	self.Head = newNode
}

func (self *LinkedList) FindTail() {
	current := self.Head
	if current.NextNode {
		current = current.NextNode
	}
	self.Tail = current
}
