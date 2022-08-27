package colections

import (
	"github.com/stretchr/testify/assert"
	"go-foundation/pkg/json"
	"testing"
)

type Test struct {
	Name string
}

var data = json.Json{"Name": "Edson"}
var data2 = Test{Name: "Edson"}

var input = []json.Json{data, data}
var input2 = []Test{data2, data2}

func TestMakeHashStruct(t *testing.T) {
	c := MakeHashStruct("Name", input2)
	assert.NotNil(t, c["Edson"])
	assert.Equal(t, input2[1], data2)
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
