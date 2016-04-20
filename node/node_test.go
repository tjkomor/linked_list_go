package node_test

import (
	. "github.com/tjkomor/linked-list/node"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	It("knows about it data", func() {
		node := Node{Data: "tyler"}
		Expect(node.Data).To(Equal("tyler"))
	})

	It("knows about next node", func() {
		node2 := Node{Data: "peter"}
		node1 := Node{Data: "tyler", NextNode: node2}
		Expect(node1.Data).To(Equal("tyler"))
		Expect(node1.NextNode).To(Equal(node2))
		Expect(node2.Data).To(Equal("peter"))
	})
})
