package main

import (
	"testing"
)

// helper function to create a new TreeNode
func newNode[T comparable](value T, children ...*TreeNode[T]) *TreeNode[T] {
	return &TreeNode[T]{value, children}
}

// TestFindShallowestDuplicate with various tree structures
func TestFindShallowestDuplicate(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode[int]
		wantValue *int
		wantDepth int
	}{
		{
			name:      "No duplicates",
			root:      newNode(1, newNode(2), newNode(3)),
			wantValue: nil,
			wantDepth: 0,
		},
		{
			name:      "Shallow duplicate",
			root:      newNode(1, newNode(1), newNode(3)),
			wantValue: new(int),
			wantDepth: 0,
		},
		{
			name:      "Deep duplicate",
			root:      newNode(1, newNode(2, newNode(3, newNode(3)))),
			wantValue: new(int),
			wantDepth: 2,
		},
		{
			name:      "Complex tree with multiple duplicates",
			root:      newNode(1, newNode(2, newNode(2)), newNode(1)),
			wantValue: new(int),
			wantDepth: 0,
		},
		{
			name:      "Single node",
			root:      newNode(1),
			wantValue: nil,
			wantDepth: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantValue != nil {
				*tc.wantValue = tc.root.value
			}

			gotValue, gotDepth := findShallowestDuplicate(tc.root)
			if (gotValue == nil || tc.wantValue == nil) && gotValue != tc.wantValue {
				t.Errorf("Test %s failed: expected nil, got non-nil or vice versa", tc.name)
			} else if gotValue != nil && tc.wantValue != nil && *gotValue != *tc.wantValue {
				t.Errorf("Test %s failed: expected %v, got %v", tc.name, *tc.wantValue, *gotValue)
			}
			if gotDepth != tc.wantDepth {
				t.Errorf("Test %s failed: expected depth %d, got %d", tc.name, tc.wantDepth, gotDepth)
			}
		})
	}
}
