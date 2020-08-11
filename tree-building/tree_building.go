// Package tree implement data structures and functions for the tree building
package tree

import "sort"

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

	// assigning child
	for _, record := range records {
		// prevent stackoverflow
		if record.ID == 0 {
			continue
		}
		if parentNode, ok := traceTable[record.Parent]; ok {
			// fmt.Println(parentNode)
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
