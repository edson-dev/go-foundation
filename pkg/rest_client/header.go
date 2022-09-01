package rest_client

type Header struct {
	Key   string
	Value string
}

func MakeHashHeader(input []Header) map[string][]string {
	mapper := make(map[string][]string)
	for _, data := range input {
		mapper[data.Key] = append(mapper[data.Key], data.Value)
	}
	return mapper
}
