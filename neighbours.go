package main

type NeighbourSet map[int]struct{}

func newNeighbourSet(list []int) NeighbourSet {
	set := make(NeighbourSet)
	for _, value := range list {
		set[value] = exists
	}
	return set
}

func (n NeighbourSet) Intersection(d NeighbourSet) NeighbourSet {
	intersection := make(NeighbourSet)
	for key, value := range n {
		if _, exists := d[key]; exists {
			intersection[key] = value
		}
	}
	return intersection
}
