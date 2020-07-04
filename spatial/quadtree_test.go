package spatial_test

import (
	"testing"

	"github.com/nvnamsss/goinf/spatial"
)

func TestRemoveItem(t *testing.T) {
	var s spatial.Quadtree
	b := spatial.Bounds{X: 10, Y: 10, Height: 1, Width: 1}
	b2 := spatial.Bounds{X: 10, Y: 10, Height: 1, Width: 1}
	s.Insert(b)
	s.Insert(b2)

	s.RemoveItem(b.Item)

	if len(s.Objects) != 1 {
		t.Error("Expected 1")
	} else {
		t.Log("Completed")
	}
}
