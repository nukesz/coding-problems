package serialization

type Node struct {
	val   string
	left  *Node
	right *Node
}

func serialize(n *Node) string {
	return ""
}

func deserialize(s string) *Node {
	return &Node{"", nil, nil}
}
