package graph_test

import (
	. "github.com/jmnarloch/gograph/graph"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Graph", func() {

	vertices := 10
	var graph Graph

	BeforeEach(func() {
		graph = New(vertices)
	})

	It("Should create empty graph with no edges", func() {

		Expect(graph.Vertices()).To(Equal(vertices))
		Expect(graph.Edges()).To(Equal(0))
	})

	It("Should add an edge to graph", func() {

		graph.AddEdge(0, 1)

		Expect(graph.Edges()).To(Equal(1))
		other := graph.Adjacent(0).Next()
		Expect(other.(int)).To(Equal(1))
		other = graph.Adjacent(1).Next()
		Expect(other.(int)).To(Equal(0))
	})
})
