package graph

import "fmt"

func NearNeighbor(G map[int]([]int), start int) (map[int]int, *int) {
	var vertexs []int
	for i := 0; i < len(G); i++ {
		vertexs = append(vertexs, i)
	}
	cost := 0
	path := make(map[int]int)
	var visited []int

	vertexLooked := start
	visited = append(visited, start)
	nextVertex := -1
	for len(visited) != len(G) {
		minVal := 256
		for _, val := range vertexs {
			if !inArray(val, visited) {
				if G[vertexLooked][val] < minVal && G[vertexLooked][val] != 0 {
					minVal = G[vertexLooked][val]
					nextVertex = val
				}
			}
		}
		path[vertexLooked+1] = nextVertex + 1
		cost += G[vertexLooked][nextVertex]
		vertexLooked = nextVertex + 1
		visited = append(visited, nextVertex)
	}

	fmt.Print(path, "\n")
	if G[visited[len(visited)-1]][start] != 0 {
		path[visited[len(visited)-1]+1] = start + 1
		cost += G[visited[len(visited)-1]][start]
		return path, &cost
	} else {
		return nil, nil
	}
}
