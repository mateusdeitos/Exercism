package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
	ParentID int
}

const ROOT_NODE_ID = 0

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var nodeMap = map[int]*Node{}
	for _, record := range records {

		if record.ID == ROOT_NODE_ID && record.Parent != ROOT_NODE_ID {
			return nil, fmt.Errorf("root node has parent")
		}

		if record.ID == record.Parent && record.ID != ROOT_NODE_ID {
			return nil, fmt.Errorf("node %d has itself as parent", record.ID)
		}

		if record.ID < record.Parent && record.ID != ROOT_NODE_ID {
			return nil, fmt.Errorf("node %d has parent with higher ID %d", record.ID, record.Parent)
		}

		if record.ID < ROOT_NODE_ID {
			return nil, fmt.Errorf("node %d has negative ID", record.ID)
		}

		if record.Parent < ROOT_NODE_ID {
			return nil, fmt.Errorf("node %d has negative parent ID", record.ID)
		}

		if record.ID >= len(records) {
			return nil, fmt.Errorf("node %d has ID higher than number of nodes %d", record.ID, len(records))
		}

		if record.Parent >= len(records) {
			return nil, fmt.Errorf("node %d has parent with ID higher than number of nodes %d", record.ID, len(records))
		}

		if _, ok := nodeMap[record.ID]; ok {
			return nil, fmt.Errorf("node %d already exists", record.ID)
		}

		nodeMap[record.ID] = &Node{ID: record.ID, ParentID: record.Parent, Children: []*Node{}}
	}

	if _, ok := nodeMap[0]; !ok {
		return nil, fmt.Errorf("root node not found")
	}

	for id := 1; id < len(nodeMap); id++ {
		node := nodeMap[id]
		parentNode := nodeMap[node.ParentID]
		parentNode.Children = append(parentNode.Children, node)
	}

	return nodeMap[0], nil
}
