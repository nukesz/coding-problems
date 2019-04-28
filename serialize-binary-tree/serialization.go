package serialization

import (
	"fmt"
	"regexp"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

func serialize(n *Node) string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("[ val: \"%s\", left: %s, right: %s ]", n.val, serialize(n.left), serialize(n.right))
}

func deserialize(s string) *Node {
	if s == "" || s == "nil" {
		return nil
	}

	if match, result := matchNode(`^\[ val: \"(\S*)\", left: nil, right: nil \]$`, s); match {
		return &Node{result[1], nil, nil}
	}

	if match, result := matchNode(`^\[ val: \"(\S*)\", left: \[(.*)\], right: nil \]$`, s); match {
		return &Node{result[1], deserialize(fmt.Sprintf("[%s]", result[2])), nil}
	}

	if match, result := matchNode(`^\[ val: \"(\S*)\", left: nil, right: \[(.*)\] \]$`, s); match {
		return &Node{result[1], nil, deserialize(fmt.Sprintf("[%s]", result[2]))}
	}

	if match, result := matchNode(`^\[ val: \"(\S*)\", left: \[(.*)\], right: \[(.*)\] \]$`, s); match {
		return &Node{result[1], deserialize(fmt.Sprintf("[%s]", result[2])), deserialize(fmt.Sprintf("[%s]", result[3]))}
	}

	return nil
}

func matchNode(pattern, node string) (bool, []string) {
	r := regexp.MustCompile(pattern)
	if r.MatchString(node) {
		return true, r.FindStringSubmatch(node)
	}
	return false, nil
}
