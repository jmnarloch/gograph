package dfs_test

import (
	. "github.com/jmnarloch/gograph/dfs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/gograph/graph"
)

var _ = Describe("DFS", func() {

	vertices := 10
	var g graph.Graph
	var dfs DFS

	BeforeEach(func() {
		g = graph.New(vertices)
		dfs = Search(g, 0)
	})

	It("should perform bread first search", func() {

		Expect(dfs).NotTo(BeNil())
	})
})
