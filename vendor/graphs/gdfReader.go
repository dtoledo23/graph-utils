package graphs

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

// ReadGDF reads a gdf file to build a Graph object.
func ReadGDF(filename string) (*Graph, error) {
	log.Println("Reading new grapf from file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	graph := New()

	reader := bufio.NewReader(file)
	lineBytes, _, err := reader.ReadLine()
	line := string(lineBytes[:])

	if strings.HasPrefix(line, nodedef) {
		nodeKeys := strings.TrimPrefix(line, nodedef)

		for _, nodeKey := range strings.Split(nodeKeys, ",") {
			graph.AddNodeKey(nodeKey)
		}

		err := readNodes(reader, graph)
		if err != nil {
			return nil, err
		}

		err = readEdges(reader, graph)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("Not valid gdf file")
	}

	return graph, nil
}

func readNodes(reader *bufio.Reader, graph *Graph) error {

	for {
		startBytes, err := reader.Peek(len(edgedef))
		lineStart := string(startBytes[:])
		if err != nil {
			return err
		}

		if strings.HasPrefix(lineStart, edgedef) {
			return nil
		}

		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			return err
		}

		line := string(lineBytes[:])
		err = readNodeFromLine(line, graph)
		if err != nil {
			return err
		}
	}
}

func readNodeFromLine(line string, graph *Graph) error {
	var id string
	properties := make([]string, 0, 10)

	for i, key := range strings.Split(line, ",") {
		if i == 0 {
			id = key
		} else {
			properties = append(properties, key)
		}
	}

	log.Println("Adding new node:", id)
	return graph.AddNode(NewNode(id, properties))
}

func readEdges(reader *bufio.Reader, graph *Graph) error {
	lineBytes, _, err := reader.ReadLine()
	if err != nil {
		return err
	}

	line := string(lineBytes[:])
	if !strings.HasPrefix(line, edgedef) {
		return errors.New("Invalid gdf format")
	}

	for {
		lineBytes, _, _ = reader.ReadLine()
		if lineBytes == nil {
			return nil
		}
		line := string(lineBytes[:])
		err := readEdgeFromLine(line, graph)
		if err != nil {
			return err
		}
	}
}

func readEdgeFromLine(line string, graph *Graph) error {
	keys := strings.Split(line, ",")
	log.Println("Adding new edge from:", keys[0], " to ", keys[1])
	return graph.AddEdge(keys[0], keys[1])
}
