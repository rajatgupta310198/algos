package graph

import (
	"testing"
)

func TestNewGraph(t *testing.T) {

	G := NewGraph()

	v0 := G.AddVertex("A")
	v1 := G.AddVertex("B")
	v2 := G.AddVertex("C")
	v3 := G.AddVertex("D")
	G.AddEdge(v0, v1)
	G.AddEdge(v0, v2)
	G.AddEdge(v1, v2)
	G.AddEdge(v2, v0)
	G.AddEdge(v2, v3)
	G.AddEdge(v3, v3)
	G.PrintGraph()
	G.Dfs(v2)

}
