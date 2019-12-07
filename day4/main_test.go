package main

import "testing"

func TestNumPossiblePasswords(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"111111-111111", 1},
		{"111111-111119", 9},
		{"111111-111122", 10},
		{"223450-223450", 0},
		{"123789-123789", 0},
	}

	for _, table := range tables {
		result := NumPossiblePasswords(table.input)
		if result != table.expected {
			t.Errorf("NumPossiblePasswords returned unexpected result, got %d instead of %d", result, table.expected)
		}
	}
}
