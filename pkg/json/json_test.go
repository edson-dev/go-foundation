package json

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	data := `{"Name": "India"}`
	a, _ := ImportString(data)
	b, _ := ImportString(data)
	fmt.Println("Value:", a, "-", b)
	fmt.Println("Address:", &a, "-", &b)
	assert.Equal(t, a, b)
	assert.True(t, &a != &b)
	//this compare the object not the pointers
	assert.Equal(t, &a, &b)
}
func TestImportExport(t *testing.T) {
	data := `{"Name": "India"}`
	a, _ := ImportString(data)
	b, _ := a.ExportString()
	require.JSONEq(t, data, b)
}
func TestImportExportAdd(t *testing.T) {
	data := `{"Name": "India"}`
	a, _ := ImportString(data)
	a.Add("test", 123)
	b, _ := a.ExportString()
	data = `{"Name": "India","test":123}`
	a, _ = ImportString(data)
	require.JSONEq(t, data, b)
}
func TestImportExportRemove(t *testing.T) {
	data := `{"Name": "India","test":123}`
	a, _ := ImportString(data)
	a.Remove("test")
	b, _ := a.ExportString()
	exp := `{"Name": "India"}`
	require.JSONEq(t, exp, b)
}
func TestImportFail(t *testing.T) {
	data := `{"Name": "India"`
	_, e := ImportString(data)

	require.EqualError(t, e, "unexpected end of JSON input")
}
