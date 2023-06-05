package orderedset

import "testing"

func TestSet_Push(t *testing.T) {
	var set Set
	set.Push(0)
	set.Push(1)
	set.Push(2)
	set.Push(3)
	if set.Count() != 4 {
		t.Fatal("count mismatch")
	}
}
