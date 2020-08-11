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
	traceTable := make(map[int]*Node)

	maxID := -1
	// creating all the nodes without assigning any child
	for _, record := range records {
		node := createNode(record)
		if _, ok := traceTable[record.ID]; ok {
			return nil, fmt.Errorf("Duplicated node")
		}
		if record.Parent > record.ID {
			return nil, fmt.Errorf("higher id parent of lower id")
		}
		if record.ID != 0 && record.ID == record.Parent {
			return nil, fmt.Errorf("direct loop detected")
		}
		traceTable[record.ID] = node
		if record.ID > maxID {
			maxID = record.ID
		}
	}
	if len(traceTable) > 0 && traceTable[0] == nil {
		return nil, fmt.Errorf("No root node")
	}
	if len(traceTable) != maxID+1 {
		return nil, fmt.Errorf("no continuous")
	}

	// assigning child
	for _, record := range records {
		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, fmt.Errorf("Root(node 0) can not have any parent")
			}
			continue
		}
		if parentNode, ok := traceTable[record.Parent]; ok {
			if parentNode.Children == nil {
				parentNode.Children = make([]*Node, 0)
			}
			parentNode.Children = append(parentNode.Children, traceTable[record.ID])
		}
	}

	// sort children
	for _, record := range records {
		sort.Sort(byID(traceTable[record.ID].Children))
	}

	return traceTable[0], err
}

func createNode(record Record) *Node {
	return &Node{record.ID, nil}
}

type byID []*Node

func (arr byID) Len() int           { return len(arr) }
func (arr byID) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }
func (arr byID) Less(i, j int) bool { return arr[i].ID < arr[j].ID }
