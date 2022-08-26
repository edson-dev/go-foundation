package colections

import "reflect"

func MakeHash[T comparable](field string, input []T) map[string]T {
	mapper := make(map[string]T)
	for _, data := range input {
		index := reflect.ValueOf(data).FieldByName(field)
		mapper[index.String()] = data
	}
	return mapper
}
func MakeHashList[T comparable](field string, input []T) map[string][]T {
	mapper := make(map[string][]T)
	for _, data := range input {
		index := reflect.ValueOf(data).FieldByName(field)
		mapper[index.String()] = append(mapper[index.String()], data)
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
