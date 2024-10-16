package utils

import "encoding/json"

func ConvertBytesToMap(value []byte) map[string]any {
	m := make(map[string]any)
	err := json.Unmarshal(value, &m)
	if err != nil {
		panic(err)
	}
	return m
}
