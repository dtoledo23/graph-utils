package graphs

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var nodedef = "nodedef>"
var edgedef = "edgedef>"

// WriteGDF reads a gdf file to build a Graph object.
func WriteGDF(graph *Graph, name string) error {
	log.Println("Writing graph on", name)
	filename, err := filepath.Abs(name)
	file, err := os.Create(filename)
	fileWriter := bufio.NewWriter(file)

	// Write node header.
	fileWriter.WriteString(nodedef)
	fileWriter.WriteString(strings.Join(graph.nodeKeys, ","))
	fileWriter.WriteRune('\n')

	// Write nodes.
	for _, node := range graph.nodes {
		fileWriter.WriteString(node.id + ",")
		fileWriter.WriteString(strings.Join(node.properties, ","))
		fileWriter.WriteRune('\n')
	}

	// Write edge header.
	fileWriter.WriteString("edgedef>node1 VARCHAR,node2 VARCHAR,directed BOOLEAN\n")
	for _, node := range graph.nodes {
		for _, neighbor := range node.neighbors {
			fileWriter.WriteString(strings.Join([]string{node.id, neighbor, "true\n"}, ","))
		}
	}
	fileWriter.Flush()
	// for _, property := range g
	return err
}
