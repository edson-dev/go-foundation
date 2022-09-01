package colections

import (
	"fmt"
	"github.com/edson-dev/go-foundation/pkg/json"
	"reflect"
)

func MakeHashStruct[T interface{}](field string, input []T) map[string]T {
	mapper := make(map[string]T)
	for i, data := range input {
		fmt.Println(input[i])
		index := reflect.ValueOf(data).FieldByName(field)
		mapper[index.String()] = data
	}
	return mapper
}

func MakeHashInterface[T json.Json](field string, input []T) map[string]interface{} {
	mapper := make(map[string]interface{})
	for _, data := range input {
		index := data[field].(string)
		mapper[index] = data
	}
	return mapper
}

func MakeHashList[T any](field string, input []T) map[string][]T {
	mapper := make(map[string][]T)
	for _, data := range input {
		index := reflect.ValueOf(data).FieldByName(field)
		mapper[index.String()] = append(mapper[index.String()], data)
	}
	return mapper
}
func MakeHashListInterface[T json.Json](field string, input []T) map[string][]T {
	mapper := make(map[string][]T)
	for _, data := range input {
		index := data[field].(string)
		mapper[index] = append(mapper[index], data)
	}
	return mapper
}
func MakeHashSet[T comparable](field string, input []T) map[string][]T {
	mapper := make(map[string][]T)
	for _, data := range input {
		index := reflect.ValueOf(data).FieldByName(field)
		if Contains(mapper[index.String()], data) == false {
			mapper[index.String()] = append(mapper[index.String()], data)
		}
	}
	return mapper
}

func Contains[T comparable](arr []T, x T) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}
