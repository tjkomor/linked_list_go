package linked_list_test

import (
	. "github.com/tjkomor/linked-list"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	It("will set head data to nil automatically", func() {
		// list := LinkedList{}
		// Expect(list.Head.Data).To(BeNil())
	})

	It("is able to find the tail", func() {
		list := LinkedList{}
		list.Append("tyler")
		list.Append("peter")
		Expect(list.Head.Data).To(Equal("tyler"))
		Expect(list.Tail).To(Equal("peter"))
	})
})
