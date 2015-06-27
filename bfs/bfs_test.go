package bfs_test

import (
	. "github.com/jmnarloch/gograph/bfs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/gograph/graph"
)

var _ = Describe("BFS", func() {

	vertices := 10
	var g graph.Graph
	var bfs BFS

	BeforeEach(func() {
		g = graph.New(vertices)
		bfs = Search(g, 0)
	})

	It("should perform bread first search", func() {

		Expect(bfs).NotTo(BeNil())
	})
})
