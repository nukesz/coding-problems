package serialization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeNil(t *testing.T) {
	var node *Node
	assert.Equal(t, "nil", serialize(node))
}

func TestDeserializeNil(t *testing.T) {
	s := "nil"
	var node *Node
	assert.Equal(t, node, deserialize(s))
}

func TestSerializeEmptyNode(t *testing.T) {
	node := &Node{}
	assert.Equal(t, `[ val: "", left: nil, right: nil ]`, serialize(node))
}

func TestDeserializeEmptyNode(t *testing.T) {
	assert.Equal(t, &Node{}, deserialize(`[ val: "", left: nil, right: nil ]`))
}

func TestSerializeNodeWithValue(t *testing.T) {
	node := &Node{
		val: "root",
	}
	assert.Equal(t, `[ val: "root", left: nil, right: nil ]`, serialize(node))
}

func TestDeserializeNodeWithValue(t *testing.T) {
	assert.Equal(t, &Node{val: "root"}, deserialize(`[ val: "root", left: nil, right: nil ]`))
}

func TestSerializeLeftNode(t *testing.T) {
	node := &Node{
		left: &Node{val: "leftNode"},
	}
	assert.Equal(t, `[ val: "", left: [ val: "leftNode", left: nil, right: nil ], right: nil ]`, serialize(node))
}

func TestDeserializeLeftNode(t *testing.T) {
	node := &Node{
		left: &Node{val: "leftNode"},
	}
	s := `[ val: "", left: [ val: "leftNode", left: nil, right: nil ], right: nil ]`
	assert.Equal(t, node, deserialize(s))
}

func TestSerializeRightNode(t *testing.T) {
	node := &Node{
		right: &Node{val: "rightNode"},
	}
	assert.Equal(t, `[ val: "", left: nil, right: [ val: "rightNode", left: nil, right: nil ] ]`, serialize(node))
}

func TestDeserializeRightNode(t *testing.T) {
	node := &Node{
		right: &Node{val: "rightNode"},
	}
	s := `[ val: "", left: nil, right: [ val: "rightNode", left: nil, right: nil ] ]`
	assert.Equal(t, node, deserialize(s))
}

func TestDeserializeLeftAndRightNode(t *testing.T) {
	node := &Node{
		left:  &Node{val: "leftNode"},
		right: &Node{val: "rightNode"},
	}
	s := `[ val: "", left: [ val: "leftNode", left: nil, right: nil ], right: [ val: "rightNode", left: nil, right: nil ] ]`
	assert.Equal(t, node, deserialize(s))
}

func TestSerializeAndDeserializeBack(t *testing.T) {
	node := &Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}
	assert.Equal(t, "left.left", deserialize(serialize(node)).left.left.val)
}
