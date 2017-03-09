package main

import (
	"bufio"
	"fmt"
	"graphs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if (len(os.Args) == 1) && (strings.HasSuffix(os.Args[1], ".gdf")) {
		fmt.Println("Usage: go run main.go <file.gdf>")
		os.Exit(0)
	}

	file := os.Args[1]
	filename, _ := filepath.Abs(file)
	fmt.Println("Reading:", filename)

	graph, err := graphs.ReadGDF(filename)

	if err != nil {
		log.Fatal(err)
	}

	ranks, err := graph.PageRank()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("results.csv")
	w := bufio.NewWriter(f)
	avg := float64(0)
	len := 0
	for k, v := range ranks {
		avg += v
		len++
		fmt.Println(k, v)
		w.WriteString(k + ",")
		w.WriteString(strconv.FormatFloat(v, 'f', 6, 64) + "\n")
	}

	fmt.Println("Average: ", avg/float64(len))

	// dfs, err := graph.DFS(origin)
	// if err != nil {
	// 	log.Fatal("Error doing DFS", err)
	// }

	// bfs, err := graph.BFS(origin)
	// if err != nil {
	// 	log.Fatal("Error doing DFS", err)
	// }

	// log.Println("Original graph:")
	// graph.Print()

	// log.Println("DFS result graph:")
	// dfs.Print()

	// log.Println("BFS result graph:")
	// bfs.Print()

	// graphs.WriteGDF(dfs, "dfs.gdf")
	// graphs.WriteGDF(bfs, "bfs.gdf")

}
