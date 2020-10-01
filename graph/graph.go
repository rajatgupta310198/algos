package graph

import (
	"fmt"
)

// Vertex Structure
type Vertex struct {
	Label   string
	AdjList []*Vertex
}

// Graph Entity
type Graph struct {
	VertexCount int
	Vertices    []*Vertex
}

// NewGraph returns
func NewGraph() *Graph {
	return new(Graph)
}

// AddVertex  adds new vertex to G of given label and returns that vertex
func (g *Graph) AddVertex(label string) *Vertex {
	vertex := new(Vertex)
	vertex.Label = label
	g.VertexCount = g.VertexCount + 1
	g.Vertices = append(g.Vertices, vertex)
	return vertex
}

// AddEdge creates an edge between two vertex
func (g *Graph) AddEdge(srcVertex, dstVertex *Vertex) *Vertex {
	srcVertex.AdjList = append(srcVertex.AdjList, dstVertex)
	return srcVertex
}

// PrintGraph prints the create graph G
func (g *Graph) PrintGraph() {
	fmt.Printf("Number of vertices %d\n", g.VertexCount)
	for _, v := range g.Vertices {
		fmt.Printf("Vertex: %s\n Adjency list \n", v.Label)
		for _, adv := range v.AdjList {
			fmt.Printf("%s, ", adv.Label)
		}
		fmt.Printf("\n")
	}
}
