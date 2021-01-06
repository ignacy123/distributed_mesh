package main

import (
	"leader/undirected_mesh"
	"lib"
	"log"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[len(os.Args)-2])
	b, _ := strconv.Atoi(os.Args[len(os.Args)-1])
	if a <= 2 || b <= 2 {
		log.Println("too small")
		return
	}
	log.Println("buliding", a, "x", b, "mesh")
	adjacencyList := make([][]int, a*b)
	for i := range adjacencyList {
		adjacencyList[i] = make([]int, 0)
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if i-1 > 0 {
				adjacencyList[(i-1)*b+j-1] = append(adjacencyList[(i-1)*b+j-1], (i-2)*b+j)
			}
			if j-1 > 0 {
				adjacencyList[(i-1)*b+j-1] = append(adjacencyList[(i-1)*b+j-1], (i-1)*b+j-1)
			}
			if i+1 <= a {
				adjacencyList[(i-1)*b+j-1] = append(adjacencyList[(i-1)*b+j-1], (i)*b+j)
			}
			if j+1 <= b {
				adjacencyList[(i-1)*b+j-1] = append(adjacencyList[(i-1)*b+j-1], (i-1)*b+j+1)
			}
		}
	}
	g, n := lib.BuildSynchronizedGraphFromAdjacencyListWithRandomIndexes(adjacencyList)
	undirected_mesh.RunMeshLeader(g, n, (a+b-2)*2)

}
