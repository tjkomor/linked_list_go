package linked_list_test

import (
	. "github.com/tjkomor/linked-list"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	It("will set head to nil automatically", func() {
		list := LinkedList{}
		Expect(list.Head).To(BeNil())
	})

	It("is able to find the tail", func() {
		list := LinkedList{}
		list.Append("tyler")
		// list.Append("peter")
		Expect(list.Head).To(Equal("tyler"))
		// Expect(list.Tail.Data).To(Equal("peter"))
	})
})
