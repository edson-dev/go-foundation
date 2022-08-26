package colections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Name string
	Data string
}

var input = []Test{{Name: "Edson", Data: "Test"}, {Name: "Edson", Data: "Test"}}

func TestMakeHash(t *testing.T) {
	c := MakeHash("Name", input)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, c["Edson"], input[1])
}
func TestMakeHashList(t *testing.T) {
	c := MakeHashList("Name", input)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, len(c["Edson"]), 2)
}
func TestMakeHashSet(t *testing.T) {
	c := MakeHashSet("Name", input)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, len(c["Edson"]), 1)
}
