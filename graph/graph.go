package graph

import (
	"github.com/jmnarloch/gograph/iterator"
	listit "github.com/jmnarloch/gograph/iterator/list"

	"container/list"
	"fmt"
)

type Graph interface {
	Vertices() int
	Edges() int
	AddEdge(v, w int)
	Adjacent(vertex int) iterator.Iterator
}

func New(vertices int) Graph {

	if vertices < 0 {
		panic("Vertices can not be negative")
	}

	graph := &graph{vertices: vertices}
	graph.adjacent = make([]*list.List, graph.vertices)
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

	g.checkBounds(v)
	g.checkBounds(w)

	g.adjacent[v].PushBack(w)
	g.adjacent[w].PushBack(v)
	g.edges++
}

func (g *graph) Adjacent(vertex int) iterator.Iterator {

	g.checkBounds(vertex)

	return listit.New(g.adjacent[vertex])
}

func (g *graph) checkBounds(vertex int) {

	if vertex < 0 || vertex >= g.vertices {
		panic(fmt.Sprintf("Vertex does exceed bounds: [%d, %d]", 0, g.vertices))
	}
}
