package main

import (
	"testing"
)

func TestRuneSetUnion(t *testing.T) {
	hello := newRuneSet("hello")
	world := newRuneSet("world")

	union := hello.Union(world).String()
	if union != "dehlorw" {
		t.Errorf("hello.Union(world) = %s; want dehlorw", union)
	}
}

func TestRuneSetIntersection(t *testing.T) {
	squiz := newRuneSet("squiz")
	chunk := newRuneSet("chunk")
	rhyme := newRuneSet("rhyme")

	squizSquiz := squiz.Intersection(squiz)
	if squizSquiz != 5 {
		t.Errorf("squiz.Intersection(squiz) = %d; want 5", squizSquiz)
	}

	squizChunk := squiz.Intersection(chunk)
	if squizChunk != 1 {
		t.Errorf("squiz.Intersection(chunk) = %d; want 1", squizChunk)
	}

	squizRhyme := squiz.Intersection(rhyme)
	if squizRhyme != 0 {
		t.Errorf("squiz.Intersection(Rhyme) = %d; want 0", squizRhyme)
	}

	chunkRhyme := chunk.Intersection(rhyme)
	if chunkRhyme != 1 {
		t.Errorf("chunk.Intersection(Rhyme) = %d; want 1", chunkRhyme)
	}
}
