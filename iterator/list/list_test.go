package list_test

import (
	. "github.com/jmnarloch/gograph/iterator/list"

	"container/list"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {

	var l *list.List
	var iterator *ListIterator

	BeforeEach(func() {
		l = list.New()
		iterator = New(l)
	})

	It("Should create list iterator", func() {

		Expect(iterator).NotTo(BeNil())
	})
})
