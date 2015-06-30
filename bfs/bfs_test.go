package bfs_test

import (
	. "github.com/jmnarloch/gograph/bfs"

	"github.com/jmnarloch/gograph/digraph"
	"github.com/jmnarloch/gograph/graph"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bread first search", func() {

	vertices := 10

	Context("Undirected graph", func() {

		var g graph.Graph
		var bfs BFS

		BeforeEach(func() {
			g = graph.New(vertices)
			for ind := 0; ind < vertices-1; ind++ {
				g.AddEdge(ind, ind+1)
			}
			bfs = Search(g, 0)
		})

		It("should perform bread first search", func() {

			Expect(bfs).NotTo(BeNil())
		})

		It("should verify that every node is connected", func() {

			for ind := 0; ind < vertices; ind++ {
				Expect(bfs.Marked(ind)).To(BeTrue())
			}
		})
	})

	Context("Directed graph", func() {

		var g digraph.Digraph
		var bfs BFS

		BeforeEach(func() {
			g = digraph.New(vertices)
			for ind := 0; ind < vertices-1; ind++ {
				g.AddEdge(ind, ind+1)
			}
			bfs = DirectedSearch(g, 0)
		})

		It("should perform bread first search", func() {

			Expect(bfs).NotTo(BeNil())
		})

		It("should verify that every node is connected", func() {

			for ind := 0; ind < vertices; ind++ {
				Expect(bfs.Marked(ind)).To(BeTrue())
			}
		})

		It("should verify that only specific nodes are connected", func() {

			for ind := 1; ind < vertices; ind++ {

				g = digraph.New(vertices)
				for i := 1; i <= ind; i++ {
					g.AddEdge(i-1, i)
				}

				bfs = DirectedSearch(g, 0)
				for j := 1; j <= ind; j++ {
					Expect(bfs.Marked(j)).To(BeTrue())
				}
				for j := ind + 1; j < vertices; j++ {
					Expect(bfs.Marked(j)).To(BeFalse())
				}
			}
		})
	})
})
