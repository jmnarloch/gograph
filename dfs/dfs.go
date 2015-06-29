package dfs

import (
	"container/list"
	"github.com/jmnarloch/gograph/digraph"
	"github.com/jmnarloch/gograph/graph"
	"fmt"
)

type DFS interface {
	Marked(vertex int) bool
}

func Search(g graph.Graph, source int) DFS {
	if(g == nil) {
		panic("Parameter 'graph' can not be nil")
	}

	dfs := &dfs{graph: g, marked: make([]bool, g.Vertices())}
	dfs.search(source)
	return dfs
}

func DirectedSearch(g digraph.Digraph, sources ...int) DFS {
	if(g == nil) {
		panic("Parameter 'digraph' can not be nil")
	}

	dfs := &dfs{graph: g, marked: make([]bool, g.Vertices())}
	for _, source := range sources {
		dfs.search(source)
	}
	return dfs
}

type dfs struct {
	graph  graph.Graph
	marked []bool
}

func (d *dfs) Marked(vertex int) bool {
	d.checkBounds(vertex)

	return d.marked[vertex]
}

func (d *dfs) search(source int) {
	d.checkBounds(source)

	if(d.marked[source]) {
		return
	}

	stack := list.New()
	stack.PushBack(source)

	for stack.Len() > 0 {
		v := stack.Remove(stack.Back()).(int)

		if !d.marked[v] {
			d.marked[v] = true

			for it := d.graph.Adjacent(v); it.HasNext(); {
				stack.PushBack(it.Next().(int))
			}
		}
	}
}

func (d *dfs) checkBounds(vertex int) {

	if vertex < 0 || vertex >= d.graph.Vertices() {
		panic(fmt.Sprintf("Vertex does exceed bounds: [%d, %d]", 0, d.graph.Vertices()))
	}
}
