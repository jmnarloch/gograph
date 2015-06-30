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
	dfs := &dfs{
		graph:  g,
		marked: make([]bool, g.Vertices()),
	}
	dfs.search(source)
	return dfs
}

func DirectedSearch(g digraph.Digraph, sources ...int) DFS {
	dfs := &dfs{
		graph:  g,
		marked: make([]bool, g.Vertices()),
	}
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
	return d.marked[vertex]
}

func (d *dfs) search(source int) {
	if d.marked[source] {
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
