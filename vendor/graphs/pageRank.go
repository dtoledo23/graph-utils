package graphs

import "math"

const (
	diffThreshold = float64(0.00001)
	initialRank   = float64(1)
	d             = float64(0.85)
)

// PageRank calculate
func (g *Graph) PageRank() (map[string]float64, error) {

	// Initialize ranks
	ranks := make(map[string]float64)
	for nodeID := range g.nodes {
		ranks[nodeID] = initialRank
	}

	// Flag to indicate when to stop iterating.
	keepIterating := true
	// i := 1
	for keepIterating {
		keepIterating = false
		for nodeID, node := range g.nodes {
			rank := 1 - d

			for _, pointToNode := range node.pointsToMe {
				PRi := ranks[pointToNode]
				Ci := float64(len(g.nodes[pointToNode].neighbors))
				rank += d * (PRi / Ci)
			}

			if math.Abs(ranks[nodeID]-rank) >= diffThreshold {
				keepIterating = true
			}

			ranks[nodeID] = rank
		}

		// log.Println("Iteration", i, ": ", ranks)
	}

	return ranks, nil
}
