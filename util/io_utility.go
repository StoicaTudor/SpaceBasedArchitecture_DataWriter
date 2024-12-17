package util

import "encoding/json"

func DeserializeJSON[T any](data string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
