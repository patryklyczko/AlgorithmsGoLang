package graph

func Dijkstra(G map[int]([]int), w [][]int, vertexSpawn int) ([][]int, int) {
	/*
		Takes the graph G with weights w and find min spawning tree
		Returns:
			- map of connecting vertexs and cost of travel
	*/
	sum := 0
	V := getKeys(G)
	var vertexTree [][]int
	alfa := make(map[int](int))
	beta := make(map[int](int))

	for _, u := range V {
		alfa[u] = 0
		beta[u] = 256
	}
	Q := V
	beta[vertexSpawn] = 0
	Q = removeByValue(Q, vertexSpawn)
	uPrime := vertexSpawn

	for len(Q) != 0 {
		neighList := G[uPrime]
		for _, i := range Q {
			if inArray(i, neighList) {
				if w[i-1][uPrime-1] < beta[i] {
					alfa[i] = uPrime
					beta[i] = w[i-1][uPrime-1]
				}
			}
		}
		minVal := 256
		for _, v := range Q {
			if beta[v] < minVal {
				uPrime = v
				minVal = beta[v]
			}
		}
		Q = removeByValue(Q, uPrime)
		vertexTree = append(vertexTree, []int{alfa[uPrime], uPrime})
		sum += w[alfa[uPrime]-1][uPrime-1]
	}
	return vertexTree, sum
}

func getKeys(m map[int]([]int)) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func removeByValue(arr []int, val int) []int {
	for i, v := range arr {
		if v == val {
			arr = append(arr[:i], arr[i+1:]...)
			return arr
		}
	}
	return []int{}
}
