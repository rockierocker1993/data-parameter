package util

import (
	"encoding/json"
	"errors"
)

// StringToObject mengonversi string JSON ke objek target
func StringToObject(data string, v interface{}) error {
	if data == "" {
		return errors.New("input string is empty")
	}
	return json.Unmarshal([]byte(data), v)
}

// ObjectToString mengonversi objek ke string JSON
func ObjectToString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
