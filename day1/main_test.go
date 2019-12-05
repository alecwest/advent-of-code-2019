package main

import "testing"

func TestTotalFuelNeeded(t *testing.T) {
	tables := []struct {
		modules  []int
		expected int
	}{
		{[]int{12}, 2},
		{[]int{14}, 2},
		{[]int{1969}, 966},
		{[]int{100756}, 50346},
	}

	for _, table := range tables {
		result := TotalFuelNeeded(table.modules)
		if result != table.expected {
			t.Errorf("TotalFuelNeeded returned unexpected result, got %d instead of %d", result, table.expected)
		}
	}
}
