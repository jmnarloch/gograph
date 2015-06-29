package bfs
import (
	"container/list"
	"github.com/jmnarloch/gograph/graph"
	"github.com/jmnarloch/gograph/digraph"
	"fmt"
)

type BFS interface {
	Marked(vertex int) bool
}

func Search(g graph.Graph, source int) BFS {
	if(g == nil) {
		panic("Parameter 'graph' can not be nil")
	}

	dfs := &bfs{graph: g, marked: make([]bool, g.Vertices())}
	dfs.search(source)
	return dfs
}

func DirectedSearch(g digraph.Digraph, sources ...int) BFS {
	if(g == nil) {
		panic("Parameter 'digraph' can not be nil")
	}

	bfs := &bfs{graph: g, marked: make([]bool, g.Vertices())}
	for _, source := range sources {
		bfs.search(source)
	}
	return bfs
}

type bfs struct {
	graph  graph.Graph
	marked []bool
}

func (b *bfs) Marked(vertex int) bool {
	b.checkBounds(vertex)

	return b.marked[vertex]
}

func (b *bfs) search(source int) {
	b.checkBounds(source)

	if(b.marked[source]) {
		return
	}

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

func (b *bfs) checkBounds(vertex int) {

	if vertex < 0 || vertex >= b.graph.Vertices() {
		panic(fmt.Sprintf("Vertex does exceed bounds: [%d, %d]", 0, b.graph.Vertices()))
	}
}