package utils

import "encoding/json"

// UnmarshalJSON Generic JSON Unmarshal function
func UnmarshalJSON[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}
