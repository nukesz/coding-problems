package serialization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeAndDeserializeBack(t *testing.T) {
	node := &Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}
	assert.Equal(t, deserialize(serialize(node)), "left.left")
}
