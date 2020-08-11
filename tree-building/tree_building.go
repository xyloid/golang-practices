// Package tree implement data structures and functions for the tree building
package tree

import "fmt"

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
	traceTable := make(map[int]*Node)

	// creating all the nodes without assigning any child
	for _, record := range records {
		node := createNode(record)
		traceTable[record.ID] = node
	}
	fmt.Println(traceTable)
	// assigning child
	for _, record := range records {
		if parentNode, ok := traceTable[record.Parent]; ok {

			parentNode.Children = append(parentNode.Children, traceTable[record.ID])
		}
	}

	return traceTable[0], err
}

func createNode(record Record) *Node {
	return &Node{record.ID, make([]*Node, 0)}
}
