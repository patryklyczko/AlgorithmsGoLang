package graph

func DFS(G map[int]([]int), s int) []int {
	var arrHelp []int
	var S []int

	S = append(S, s)

	for len(S) != 0 {
		a := S[0]
		S = append(S[1:])
		if !inArray(a, arrHelp) {
			arrHelp = append(arrHelp, a)
			if checkForKey(G, a) {
				for _, i := range G[a] {
					if !inArray(i, arrHelp) {
						S = append(S, i)
					}
				}
			}
		}
	}
	return arrHelp
}

func checkForKey(m map[int]([]int), k int) bool {
	for key := range m {
		if key == k {
			return true
		}
	}
	return false
}

func inArray(value int, arr []int) bool {
	for _, val := range arr {
		if val == value {
			return true
		}
	}
	return false
}
