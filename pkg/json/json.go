package json

import (
	"encoding/json"
)

type Json map[string]interface{}

var Value Json

func ImportByte(data []byte) (*Json, error) {
	new := make(Json)
	err := json.Unmarshal(data, &new)
	return &new, err
}
func ImportString(data string) (*Json, error) {
	return ImportByte([]byte(data))
}
func (j Json) Add(field string, value interface{}) {
	j[field] = value
}
func (j Json) Remove(field string) {
	delete(j, field)
}
func (j *Json) Export() ([]byte, error) {
	export, err := json.Marshal(j)
	return export, err
}
func (j *Json) ExportString() (string, error) {
	v, e := j.Export()
	return string(v), e
}
