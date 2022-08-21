package main

import (
	"testing"
)

func TestNeighbourSetIntersection(t *testing.T) {
	a := newNeighbourSet([]int{1, 2, 3})
	b := newNeighbourSet([]int{3, 4, 5})

	intersection := a.Intersection(b)
	length := len(intersection)

	if length != 1 {
		t.Errorf("len(a.Intersection(b)) = %d; want 1", length)
	}

	if _, exists := intersection[3]; !exists {
		t.Errorf("Expecting intersection to contain 3")
	}
}
