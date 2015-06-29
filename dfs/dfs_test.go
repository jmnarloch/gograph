package dfs_test

import (
	. "github.com/jmnarloch/gograph/dfs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jmnarloch/gograph/graph"
	"github.com/jmnarloch/gograph/digraph"
)

var _ = Describe("Depth First Search", func() {

	vertices := 10

	Context("Undirected graph", func() {

		var g graph.Graph
		var dfs DFS

		BeforeEach(func() {
			g = graph.New(vertices)
			for ind := 0; ind < vertices - 1; ind++ {
				g.AddEdge(ind, ind + 1)
			}
			dfs = Search(g, 0)
		})

		It("should perform depth first search", func() {

			Expect(dfs).NotTo(BeNil())
		})

		It("should verify that every node is connected", func() {

			for ind := 0; ind < vertices; ind++ {
				Expect(dfs.Marked(ind)).To(BeTrue())
			}
		})
	})

	Context("Directed graph", func() {

		var g digraph.Digraph
		var dfs DFS

		BeforeEach(func() {
			g = digraph.New(vertices)
			for ind := 0; ind < vertices - 1; ind++ {
				g.AddEdge(ind, ind + 1)
			}
			dfs = DirectedSearch(g, 0)
		})

		It("should perform depth first search", func() {

			Expect(dfs).NotTo(BeNil())
		})

		It("should verify that every node is connected", func() {

			for ind := 0; ind < vertices; ind++ {
				Expect(dfs.Marked(ind)).To(BeTrue())
			}
		})

		It("should verify that only specific nodes are connected", func() {

			for ind := 1; ind < vertices; ind++ {

				g = digraph.New(vertices)
				for i := 1; i <= ind; i++ {
					g.AddEdge(i - 1, i)
				}

				dfs = DirectedSearch(g, 0)
				for j := 1; j <= ind; j++ {
					Expect(dfs.Marked(j)).To(BeTrue())
				}
				for j := ind + 1; j < vertices; j++ {
					Expect(dfs.Marked(j)).To(BeFalse())
				}
			}
		})
	})
})
