package colections

import (
	"github.com/edson-dev/go-foundation/pkg/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Name string
}

var j = json.Json{"Name": "Edson"}
var d = Test{Name: "Edson"}

var input = []json.Json{j, j}
var input2 = []Test{d, d}

func TestMakeHashStruct(t *testing.T) {
	c := MakeHashStruct("Name", input2)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, input2[1], d)
}
func TestMakeHashInterface(t *testing.T) {
	c := MakeHashInterface("Name", input)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, c["Edson"], input[1])
}
func TestMakeHashList(t *testing.T) {
	c := MakeHashList("Name", input2)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, len(c["Edson"]), 2)
}
func TestMakeHashListInterface(t *testing.T) {
	c := MakeHashListInterface("Name", input)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, len(c["Edson"]), 2)
}
func TestMakeHashSet(t *testing.T) {
	c := MakeHashSet("Name", input2)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, len(c["Edson"]), 1)
}
