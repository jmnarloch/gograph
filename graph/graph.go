package graph

import (
	"github.com/jmnarloch/gograph/iterator"
	listit "github.com/jmnarloch/gograph/iterator/list"

	"container/list"
)

type Graph interface {
	Vertices() int
	Edges() int
	AddEdge(v, w int)
	Adjacent(vertex int) iterator.Iterator
}

func New(vertices int) Graph {
	graph := &graph{
		vertices: vertices,
		adjacent: make([]*list.List, vertices),
	}
	for ind := 0; ind < graph.vertices; ind++ {
		graph.adjacent[ind] = list.New()
	}
	return graph
}

type graph struct {
	vertices int
	edges    int
	adjacent []*list.List
}

func (g *graph) Vertices() int {
	return g.vertices
}

func (g *graph) Edges() int {
	return g.edges
}

func (g *graph) AddEdge(v, w int) {
	g.adjacent[v].PushBack(w)
	g.adjacent[w].PushBack(v)
	g.edges++
}

func (g *graph) Adjacent(vertex int) iterator.Iterator {
	return listit.New(g.adjacent[vertex])
}
