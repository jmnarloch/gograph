package digraph

import (
	"container/list"
	"fmt"
	"github.com/jmnarloch/gograph/graph"
	"github.com/jmnarloch/gograph/iterator"
	listit "github.com/jmnarloch/gograph/iterator/list"
)

type Digraph interface {
	graph.Graph

	Reverse() Digraph
}

func New(vertices int) Digraph {

	if vertices < 0 {
		panic("Vertices can not be negative")
	}

	graph := &digraph{vertices: vertices}
	graph.adjacent = make([]*list.List, graph.vertices)
	for ind := 0; ind < graph.vertices; ind++ {
		graph.adjacent[ind] = list.New()
	}
	return graph
}

type digraph struct {
	vertices int
	edges    int
	adjacent []*list.List
}

func (g *digraph) Vertices() int {
	return g.vertices
}

func (g *digraph) Edges() int {
	return g.edges
}

func (g *digraph) AddEdge(v, w int) {

	g.checkBounds(v)
	g.checkBounds(w)

	g.adjacent[v].PushBack(w)
	g.edges++
}

func (g *digraph) Adjacent(vertex int) iterator.Iterator {

	g.checkBounds(vertex)

	return listit.New(g.adjacent[vertex])
}

func (g *digraph) Reverse() Digraph {

	reversed := New(g.vertices)
	for v := 0; v < g.vertices; v++ {
		for it := g.Adjacent(v); it.HasNext(); {
			reversed.AddEdge(it.Next().(int), v)
		}
	}
	return reversed
}

func (g *digraph) checkBounds(vertex int) {

	if vertex < 0 || vertex >= g.vertices {
		panic(fmt.Sprintf("Vertex does exceed bounds: [%d, %d]", 0, g.vertices))
	}
}
