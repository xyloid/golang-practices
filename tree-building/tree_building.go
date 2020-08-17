// Package tree implement data structures and functions for the tree building
package tree

import (
	"fmt"
	"sort"
)

// Record is a single piece of record
type Record struct {
	ID     int
	Parent int
}

// Node is the tree node in a tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree based on the given records
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))

	// creating all the nodes without assigning any child
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("Invalid records %+v", r)
		}

		// create
		nodes[i] = &Node{r.ID, nil}

		// add to parent's node
		if r.ID > 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[i])
		}
	}

	return nodes[0], nil
}
