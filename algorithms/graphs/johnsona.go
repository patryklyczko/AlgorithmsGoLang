package graph

type Traces struct {
	Cost map[int]int
	Path map[int]int
}

func Johnsona(G map[int]([]int), weight [][]int) []Traces {
	var traces []Traces
	V := getKeys(G)
	lengthV := len(V)

	modifiedWeight := BellmanFord(G, weight)
	if modifiedWeight == nil {
		return nil
	}

	modifiedWeightGraph := weight
	for i := 0; i < lengthV; i++ {
		for j := 0; j < lengthV; j++ {
			if weight[i][j] != 0 {
				modifiedWeightGraph[i][j] = weight[i][j] + modifiedWeight[i] - modifiedWeight[j]
			}
		}
	}
	for i := 1; i < lengthV+1; i++ {
		traceOne, traceTwo := DijkstraPath(G, modifiedWeightGraph, i)
		trace := Traces{Path: traceTwo, Cost: traceOne}
		traces = append(traces, trace)
	}

	for i := 0; i < lengthV; i++ {
		for j := 0; j < lengthV; j++ {
			(traces[i].Cost)[j] = (traces[i].Cost)[j] - modifiedWeight[i] + modifiedWeight[j]
		}
	}
	return traces
}
