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
func Build(records []Record) (root *Node, err error) {

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodeNumber := len(records)

	if nodeNumber == 0 {
		return
	}
	if records[nodeNumber-1].ID+1 != nodeNumber {
		return nil, fmt.Errorf("non-continuous")
	}

	if records[0].ID != 0 {
		return nil, fmt.Errorf("No root node")
	}

	if records[0].Parent != 0 {
		return nil, fmt.Errorf("Root(node 0) can not have any parent")
	}

	nodes := make([]*Node, nodeNumber)

	// creating all the nodes without assigning any child
	for i, record := range records {

		// check
		if nodes[i] != nil {
			return nil, fmt.Errorf("Duplicated node")
		}
		if record.Parent > record.ID {
			return nil, fmt.Errorf("higher id parent of lower id")
		}
		if record.ID != 0 && record.ID == record.Parent {
			return nil, fmt.Errorf("direct loop detected")
		}

		// create
		node := &Node{record.ID, nil}
		nodes[i] = node
	}

	// assigning child, note the order of children is already sorted
	for _, record := range records {
		if record.ID > 0 {
			parentNode := nodes[record.Parent]

			if parentNode.Children == nil {
				parentNode.Children = make([]*Node, 0)
			}
			parentNode.Children = append(parentNode.Children, nodes[record.ID])
		}

	}

	return nodes[0], err
}
