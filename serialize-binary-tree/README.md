Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

```go
type Node struct {
	val   string
	left  *Node
	right *Node
}
```

The following test should pass:

```go
	node := &Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}
	assert.Equal(t, deserialize(serialize(node)), "left.left")
```
