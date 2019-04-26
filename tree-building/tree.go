package tree

import "fmt"

// Node construction tree
type Node struct {
	Children []*Node
	ID       int
}

// Record for input
type Record struct {
	ID     int
	Parent int
}

func createNode(value int) *Node {
	new := Node{
		ID: value,
	}
	return &new
}

func traverseNode(parentVal int, parent *Node) *Node {
	currentVal := (*parent).ID // or simply parent.ID. Why?
	if currentVal == parentVal {
		return parent
	}
	if parent.Children != nil {
		return traverseNode(parentVal, parent.Children[0])
	}
	return nil
}

// Build to build the function
func Build(records []Record) (*Node, error) {
	fmt.Printf("%v", records)
	var parent *Node

	for _, record := range records {
		if parent == nil {
			parent = &Node{
				ID: 0,
			}
		}
		fmt.Printf("%v", record)
		if record.ID == 0 {
			continue
		}
		parentVal := record.Parent
		parentNode := traverseNode(parentVal, parent)
		(*parentNode).Children = append((*parentNode).Children, createNode(record.ID))
	}

	return parent, nil
}
