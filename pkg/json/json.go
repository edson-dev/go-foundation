package json

import (
	"encoding/json"
	"fmt"
)

type Json map[string]interface{}

func ImportByte(data []byte) Json {
	new := make(Json)
	err := json.Unmarshal(data, &new)
	if err != nil {
		fmt.Printf("Error Inporting data json")
	}
	return new
}
func ImportString(data string) Json {
	return ImportByte([]byte(data))
}
func (j *Json) Export() []byte {
	export, err := json.Marshal(j)
	if err != nil {
		fmt.Printf("Error Exporting data json")
	}
	return export
}
func (j *Json) ExportString() string {
	return string(j.Export())
}
