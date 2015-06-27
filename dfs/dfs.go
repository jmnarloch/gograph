package dfs

import (
	"container/list"
	"github.com/jmnarloch/gograph/digraph"
	"github.com/jmnarloch/gograph/graph"
)

type DFS interface {
	Marked(vertex int) bool
}

func Search(g graph.Graph, source int) DFS {
	// TODO validate the input

	dfs := &dfs{graph: g, marked: make([]bool, g.Vertices())}
	dfs.search(source)
	return dfs
}

func DirectedSearch(g digraph.Digraph, sources ...int) DFS {
	// TODO validate the input

	dfs := &dfs{graph: g, marked: make([]bool, g.Vertices())}
	for source := range sources {
		dfs.search(source)
	}
	return dfs
}

type dfs struct {
	graph  graph.Graph
	marked []bool
}

func (d *dfs) Marked(vertex int) bool {
	// TODO validate the input

	return d.marked[vertex]
}

func (d *dfs) search(source int) {

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
