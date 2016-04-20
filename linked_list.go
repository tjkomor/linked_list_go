package linked_list

import "github.com/tjkomor/linked-list/node"

type LinkedList struct {
	Head node.Noder
}

func (self LinkedList) Append(data string) {
	new_node := &node.Node{Data: data}
	println(new_node.Data)
	self.Head = new_node
}
