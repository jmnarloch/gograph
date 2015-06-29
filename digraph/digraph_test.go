package digraph_test

import (
	. "github.com/jmnarloch/gograph/digraph"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Digraph", func() {

	vertices := 10
	var digraph Digraph

	BeforeEach(func() {
		digraph = New(vertices)
	})

	It("Should create empty graph with no edges", func() {

		Expect(digraph.Vertices()).To(Equal(vertices))
		Expect(digraph.Edges()).To(Equal(0))
	})

	It("Should add an edge to graph", func() {

		digraph.AddEdge(0, 1)

		Expect(digraph.Edges()).To(Equal(1))
		other := digraph.Adjacent(0).Next()
		Expect(other.(int)).To(Equal(1))
		Expect(digraph.Adjacent(1).HasNext()).To(BeFalse())
	})

	It("Should iterate over all adjecent edges", func() {

		digraph.AddEdge(0, 1)
		digraph.AddEdge(0, 2)
		digraph.AddEdge(0, 3)

		val := 1
		for it := digraph.Adjacent(0); it.HasNext(); {
			Expect(it.Next()).To(Equal(val))
			val++
		}
		// additionally verify that the 'reverse' path exists
		for ind := 1; ind <= 3; ind++ {
			Expect(digraph.Adjacent(ind).HasNext()).To(BeFalse())
		}
	})
})
