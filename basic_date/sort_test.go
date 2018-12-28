package basic_date

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	x := 11
	s := []int{3, 6, 8, 11, 45}
	pos := sort.Search(len(s), func(i int) bool {
		return s[i] == x
	})

	if pos < len(s) && s[pos] == x {
		t.Error("baohan")
	} else {
		t.Error("no find")
	}
}
