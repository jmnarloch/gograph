package digraph

import (
	"container/list"
	"github.com/jmnarloch/gograph/graph"
	"github.com/jmnarloch/gograph/iterator"
	listit "github.com/jmnarloch/gograph/iterator/list"
)

type Digraph interface {
	graph.Graph

	Reverse() Digraph
}

func New(vertices int) Digraph {
	graph := &digraph{
		vertices: vertices,
		adjacent: make([]*list.List, vertices),
	}
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
	g.adjacent[v].PushBack(w)
	g.edges++
}

func (g *digraph) Adjacent(vertex int) iterator.Iterator {
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
