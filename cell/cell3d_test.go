package cell

import (
	"testing"
)

func TestCell3D(t *testing.T) {
	c := CreateCell3D(0, 0, 0)
	if c.links == nil {
		t.Error("links should not be nil")
	}

	cellLinks := []Cell{}
	for key, _ := range c.links {
		cellLinks = append(cellLinks, key)
	}
}
