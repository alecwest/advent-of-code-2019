package main

import "testing"

func TestNumPossiblePasswords(t *testing.T) {
	tables := []struct {
		input    string
		expected int
	}{
		{"111111-111111", 0},
		{"112233-112233", 1},
		{"123444-123444", 0},
		{"111111-111119", 0},
		{"111111-111122", 1},
		{"111122-111122", 1},
		{"223450-223450", 0},
		{"123789-123789", 0},
		{"000000-000100", 9},
		{"000000-000200", 25},
	}

	for _, table := range tables {
		result := NumPossiblePasswords(table.input)
		if result != table.expected {
			t.Errorf("NumPossiblePasswords returned unexpected result for %s, got %d instead of %d", table.input, result, table.expected)
		}
	}
}
