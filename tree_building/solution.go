package tree

import (
	"errors"
	"fmt"
	"sort"
)

/* -------------------------------------------------------------------------- */

func removeFromArray(rs []Record, r Record) []Record {
	result := []Record{}
	for _, record := range rs {
		if record != r {
			result = append(result, record)
		}
	}
	return result
}

func Contains(rs []Record, r Record, rootSeen bool) bool {
	rootRecord := Record{ID: 0, Parent: 0}
	if !rootSeen && r == rootRecord {
		return false
	}
	for _, n := range rs {
		if r == n {
			return true
		}
	}
	return false
}

/* -------------------------------------------------------------------------- */

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

/* -------------------------------------------------------------------------- */

func findAllChildren(records []Record, ID int) []*Node {
	var result []*Node
	recordsWithoutRoot := records
	for _, r := range records {
		if r.Parent == ID {
			recordsWithoutRoot = removeFromArray(recordsWithoutRoot, r)
			result = append(result, &Node{ID: r.ID, Children: findAllChildren(recordsWithoutRoot, r.ID)})
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result

}

func countLen(nodes []*Node) int {
	result := 1
	for _, n := range nodes {
		if len(n.Children) != 0 {
			result += countLen(n.Children)
		} else {
			result++
		}
	}
	return result
}

/* -------------------------------------------------------------------------- */

func Build(records []Record) (*Node, error) {
	var root Record
	recordsWithoutRoot := records
	var seenRecords []Record
	isRootInside := false

	if len(records) == 0 {
		return nil, nil
	}

	for _, r := range records {
		if Contains(seenRecords, r, isRootInside) {
			return nil, errors.New("Duplicated record")
		}
		if r.ID >= len(records) {
			return nil, errors.New("Non continuous records")
		}
		if r.ID < r.Parent {
			return nil, errors.New("Invalid node")
		}
		if r.ID == r.Parent {
			recordsWithoutRoot = removeFromArray(recordsWithoutRoot, r)
			root = r
			isRootInside = true
			fmt.Println("CALLED")
		}
		seenRecords = append(seenRecords, r)
	}
	if !isRootInside {
		return nil, errors.New("No root")
	}

	result := &Node{ID: root.ID, Children: findAllChildren(recordsWithoutRoot, root.ID)}

	if countLen(result.Children) != len(records) {
		return nil, errors.New("Non cycled record")
	}

	return result, nil
}
