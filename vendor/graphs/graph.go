package graphs

import "errors"
import "fmt"

var defaultSize = 50

// Graph represents a graph.
type Graph struct {
	size     int
	nodes    map[string]*Node
	nodeKeys []string
}

// Node represents a Node in the Graph object.
type Node struct {
	id         string
	properties []string
	neighbors  []string
}

// New initializes a new graph.
func New() *Graph {
	return &Graph{
		size:  0,
		nodes: make(map[string]*Node),
	}
}

// NewNode initializes a new node.
func NewNode(id string, properties []string) *Node {
	return &Node{
		id:         id,
		properties: properties,
		neighbors:  make([]string, 0, 10),
	}
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(n *Node) error {
	if _, ok := g.nodes[n.id]; ok {
		return errors.New("Node already exist")
	}
	g.nodes[n.id] = n
	g.size++
	return nil
}

// AddEdge add a new edge from the source node to the destination node.
func (g *Graph) AddEdge(src, dest string) error {
	_, existSrc := g.nodes[src]
	_, existDest := g.nodes[dest]

	if (!existSrc) || (!existDest) {
		return errors.New("Non-existing nodes provided")
	}

	g.nodes[src].neighbors = append(g.nodes[src].neighbors, dest)

	return nil
}

// AddNodeKey adds a new node key.
func (g *Graph) AddNodeKey(key string) {
	g.nodeKeys = append(g.nodeKeys, key)
}

// Print prints nicely a graph.
func (g *Graph) Print() {
	fmt.Println("Graph size: ", g.size)
	for _, v := range g.nodes {
		v.print()
	}
}

func (n *Node) print() {
	fmt.Println("Node:", n.id, len(n.neighbors), n.neighbors)
}

func (g *Graph) partialClone() *Graph {
	clone := New()
	for _, key := range g.nodeKeys {
		clone.AddNodeKey(key)
	}

	return clone
}
