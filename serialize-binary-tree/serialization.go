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

	r := regexp.MustCompile(`^\[ val: \"(\S*)\", left: nil, right: nil \]$`)
	result := r.FindStringSubmatch(s)
	if r.MatchString(s) {
		return &Node{result[1], nil, nil}
	}

	r = regexp.MustCompile(`^\[ val: \"(\S*)\", left: \[(.*)\], right: nil \]$`)
	result = r.FindStringSubmatch(s)
	if r.MatchString(s) {
		return &Node{result[1], deserialize(fmt.Sprintf("[%s]", result[2])), nil}
	}

	r = regexp.MustCompile(`^\[ val: \"(\S*)\", left: nil, right: \[(.*)\] \]$`)
	result = r.FindStringSubmatch(s)
	if r.MatchString(s) {
		return &Node{result[1], nil, deserialize(fmt.Sprintf("[%s]", result[2]))}
	}

	r = regexp.MustCompile(`^\[ val: \"(\S*)\", left: \[(.*)\], right: \[(.*)\] \]$`)
	result = r.FindStringSubmatch(s)
	if r.MatchString(s) {
		return &Node{result[1], deserialize(fmt.Sprintf("[%s]", result[2])), deserialize(fmt.Sprintf("[%s]", result[3]))}
	}

	return nil
}
