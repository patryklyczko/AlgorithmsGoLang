package graph

func BellmanFord(G map[int]([]int), weights [][]int) map[int]int {
	/*
		Classic Bellman Ford algorithm
	*/

	path := make(map[int](int))
	V := getKeys(G)
	lengthV := len(V)
	for i := 0; i < lengthV+1; i++ {
		path[i] = 256
	}
	path[lengthV] = 0
	var val [][]int
	for i := 0; i < lengthV; i++ {
		for j := 0; j < lengthV; j++ {
			if weights[i][j] != 256 && weights[i][j] != 0 {
				val = append(val, []int{i, j, weights[i][j]})
			}
		}
	}
	for i := 0; i < lengthV; i++ {
		val = append(val, []int{lengthV, i, 0})
	}

	for i := 0; i < lengthV; i++ {
		for j := 0; j < len(val); j++ {
			if path[val[j][0]] != 256 && path[val[j][0]]+val[j][2] < path[val[j][1]] {
				path[val[j][1]] = path[val[j][0]] + val[j][2]
			}
		}
	}
	delete(path, lengthV)
	return path
}
