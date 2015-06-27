package bfs
import (
	"container/list"
	"github.com/jmnarloch/gograph/graph"
	"github.com/jmnarloch/gograph/digraph"
)

type BFS interface {
	Marked(vertex int) bool
}

func Search(g graph.Graph, source int) BFS {
	// TODO validate the input

	dfs := &bfs{graph: g, marked: make([]bool, g.Vertices())}
	dfs.search(source)
	return dfs
}

func DirectedSearch(g digraph.Digraph, sources ...int) BFS {
	// TODO validate the input

	bfs := &bfs{graph: g, marked: make([]bool, g.Vertices())}
	for source := range sources {
		if(!bfs.marked[source]) {
			bfs.search(source)
		}
	}
	return bfs
}

type bfs struct {
	graph  graph.Graph
	marked []bool
}

func (b *bfs) Marked(vertex int) bool {
	// TODO validate the input

	return b.marked[vertex]
}

func (b *bfs) search(source int) {

	stack := list.New()
	stack.PushBack(source)

	for stack.Len() > 0 {
		v := stack.Remove(stack.Front()).(int)

		if !b.marked[v] {
			b.marked[v] = true

			for it := b.graph.Adjacent(v); it.HasNext(); {
				stack.PushBack(it.Next())
			}
		}
	}
}