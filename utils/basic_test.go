package utils

import (
	"testing"
)

func TestSort(t *testing.T) {
	items := []string{"1", "2", "5", "4"}
	Sort(items)
	if items[2] != "4" {
		t.Error("Expected string", "4", "got", items[2])
	}

	items = []string{}
	Sort(items)
}
