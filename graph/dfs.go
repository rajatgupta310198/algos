package graph

import (
	"fmt"
)

// Dfs performs Depth Fist Search
func (g *Graph) Dfs(vertex *Vertex) {
	//visited := []bool{false, false}

	visited := make(map[string]bool)
	for i := range g.Vertices {
		visited[g.Vertices[i].Label] = false
	}
	g.dfsUtil(vertex, visited)

}

// dfsUtil is utility for Dfs
func (g *Graph) dfsUtil(vertex *Vertex, visited map[string]bool) {

	visited[vertex.Label] = true
	fmt.Printf("\n%s ", vertex.Label)
	for _, v := range vertex.AdjList {
		if visited[v.Label] == false {
			g.dfsUtil(v, visited)
		}
	}

}
