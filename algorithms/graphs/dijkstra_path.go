package graph

func DijkstraPath(G map[int]([]int), weight [][]int, start int) (map[int]int, map[int]int) {
	V := getKeys(G)

	d := make(map[int]int)
	p := make(map[int]int)

	for _, val := range V {
		d[val] = 256
		p[val] = 0
	}

	Q := V
	Q = removeByValue(Q, start)
	d[start] = 0
	uPrime := start

	for len(Q) != 0 {
		for _, val := range Q {
			N := G[uPrime]
			if inArray(val, N) {
				if d[uPrime]+weight[uPrime-1][val-1] < d[val] {
					d[val] = d[uPrime] + weight[uPrime-1][val-1]
					p[val] = uPrime
				}
			}
		}
		min_val := 256
		for _, val := range Q {
			if d[val] < min_val {
				min_val = d[val]
				uPrime = val
			}
		}
		Q = removeByValue(Q, uPrime)
	}
	return d, p
}
