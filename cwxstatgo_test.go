package cwxstatgo

import (
	"testing"
)

// Quick general test
func TestCwxstatgo(t *testing.T) {
	var h *Tree
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h = Insert(h, a)
	WalkerRun(h)

	if Add(h) != 55 {
		t.Error("Add(h): Expected 55, got ", Add(h))
	}
	if Nodes(h) != 6 {
		t.Error("Nodes(h): Expected 6, got ", Nodes(h))
	}

}
