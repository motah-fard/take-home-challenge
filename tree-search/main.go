package main

import (
	"fmt"
	"math"
)

// TreeNode is now a generic type, where T is the type of the value.
type TreeNode[T comparable] struct {
	value    T
	children []*TreeNode[T]
}

func findShallowestDuplicate[T comparable](root *TreeNode[T]) (*T, int) {
	if root == nil {
		return nil, 0
	}

	queue := []*TreeNode[T]{root}
	depths := []*int{new(int)}
	depthMap := make(map[T]int)
	var shallowestDuplicate *T
	shallowestDepth := math.MaxInt32

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := *depths[i]

		if storedDepth, ok := depthMap[node.value]; ok {
			if storedDepth < shallowestDepth {
				shallowestDepth = storedDepth
				duplicate := node.value // create a copy of node.value
				shallowestDuplicate = &duplicate
			}
		} else {
			depthMap[node.value] = depth
		}

		for _, child := range node.children {
			queue = append(queue, child)
			newDepth := depth + 1
			depths = append(depths, &newDepth)
		}
	}

	if shallowestDuplicate == nil {
		return nil, 0 // No duplicate found
	}

	return shallowestDuplicate, shallowestDepth
}

func main() {
	// Example usage with int and string
	rootInt := &TreeNode[int]{1, []*TreeNode[int]{
		{2, []*TreeNode[int]{{2, nil}}},
		{3, nil},
	}}

	rootString := &TreeNode[string]{"Kiwi", []*TreeNode[string]{
		{"banana", []*TreeNode[string]{
			{"banana", []*TreeNode[string]{
				{"banana", nil},
			}},
			{"cherry", nil},
		}},
		{"date", []*TreeNode[string]{
			{"fig", []*TreeNode[string]{
				{"grape", nil},
			}},
			{"apple", []*TreeNode[string]{
				{"fig", []*TreeNode[string]{
					{"grape", nil},
				}},
				{"Kiwi", nil},
			}},
		}},
	}}

	duplicateInt, depthInt := findShallowestDuplicate(rootInt)
	if duplicateInt != nil {
		fmt.Printf("The shallowest duplicate int node is %d at depth %d\n", *duplicateInt, depthInt)
	} else {
		fmt.Println("No int duplicates found")
	}

	duplicateString, depthString := findShallowestDuplicate(rootString)
	if duplicateString != nil {
		fmt.Printf("The shallowest duplicate string node is %s at depth %d\n", *duplicateString, depthString)
	} else {
		fmt.Println("No string duplicates found")
	}
}
