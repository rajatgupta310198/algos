package graph

import (
	"fmt"
	"github.com/rajatgupta310198/algos/queue"
)

// BfsTraverse function to just traverse graph from 0th node demo purpose
func (g *Graph) BfsTraverse() {
	myGraphQueue := queue.NewQueue()
	visitedNode := make(map[string]bool)
	for i := range g.Vertices {
		visitedNode[g.Vertices[i].Label] = false
	}
	myGraphQueue.Add(g.Vertices[0])
	visitedNode[g.Vertices[0].Label] = true
	for {
		node, err := myGraphQueue.Pop()
		assertedNode, ok := node.(*Vertex)
		fmt.Println(assertedNode.Label)
		visitedNode[assertedNode.Label] = true
		if err != nil || ok == false {
			break
		}
		for _, v := range assertedNode.AdjList {
			if visitedNode[v.Label] != true {
				myGraphQueue.Add(v)
			}
		}
		if myGraphQueue.IsEmpty() == true {
			break
		}
	}
	fmt.Printf("\n")
}
