package list_test

import (
	. "github.com/jmnarloch/gograph/iterator/list"

	"container/list"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {

	count := 10
	var l *list.List
	var iterator *ListIterator

	BeforeEach(func() {
		l = list.New()

		for num := 0; num < count; num++ {
			l.PushBack(num)
		}

		iterator = New(l)
	})

	It("Should create list iterator", func() {

		Expect(iterator).NotTo(BeNil())
	})

	It("Should iterate over list elements", func() {

		for val := 0; iterator.HasNext(); val++ {
			Expect(iterator.Next()).To(Equal(val))
		}
	})
})
