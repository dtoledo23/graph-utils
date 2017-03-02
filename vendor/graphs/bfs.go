package graphs

import "log"

/*
Breadth-First-Search(Graph, root):

    create empty set S
    create empty queue Q

    root.parent = NIL
    add root to S
    Q.enqueue(root)

    while Q is not empty:
        current = Q.dequeue()
        if current is the goal:
            return current
        for each node n that is adjacent to current:
            if n is not in S:
                add n to S
                n.parent = current
                Q.enqueue(n)
*/

// BFS performs Breadth First Search algorithm.
func (g *Graph) BFS(id string) (*Graph, error) {
	log.Println("BFS started on ", id)

	// Initialize objects
	bfsGraph := g.partialClone()
	q := newQueue()
	visitedNodes := make(map[string]bool)

	// BFS
	q.Push(id)

	for !q.IsEmpty() {
		currentID, err := q.Pop()
		// Building result graph.
		bfsGraph.AddNode(g.nodes[currentID])

		if err != nil {
			return nil, err
		}

		log.Println("BFS: Visiting", currentID)
		if _, visited := visitedNodes[currentID]; !visited {
			visitedNodes[currentID] = true
			for _, n := range g.nodes[currentID].neighbors {
				q.Push(n)
				bfsGraph.AddNode(g.nodes[n])
				bfsGraph.AddEdge(currentID, n)
			}
		}
	}

	log.Println("BFS succesfully finished on ", id)
	return bfsGraph, nil
}
