package paths

type Adjacency struct {
	neighbor int
	weight   int
}

type AdjacencyList [][]Adjacency
