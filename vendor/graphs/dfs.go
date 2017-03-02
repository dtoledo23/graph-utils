package graphs

import "log"

/*
1  procedure DFS-iterative(G,v):
2      let S be a stack
3      S.push(v)
4      while S is not empty
5          v = S.pop()
6          if v is not labeled as discovered:
7              label v as discovered
8              for all edges from v to w in G.adjacentEdges(v) do
9                  S.push(w)
*/

// DFS performs Depth First Search algorithm.
func (g *Graph) DFS(id string) (*Graph, error) {
	log.Println("DFS started on ", id)

	// Create a graph clone
	dfsGraph := g.partialClone()

	s := newStack()
	visitedNodes := make(map[string]bool)

	s.Push(id)

	for !s.IsEmpty() {
		currentID, err := s.Pop()
		dfsGraph.AddNode(g.nodes[currentID])

		if err != nil {
			return nil, err
		}

		log.Println("DFS: Visiting", currentID)
		if _, visited := visitedNodes[currentID]; !visited {
			visitedNodes[currentID] = true
			for _, n := range g.nodes[currentID].neighbors {
				s.Push(n)
				dfsGraph.AddNode(g.nodes[n])
				dfsGraph.AddEdge(currentID, n)
			}
		}
	}

	log.Println("DFS succesfully finished on", id)
	return dfsGraph, nil
}
