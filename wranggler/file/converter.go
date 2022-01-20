package file

import (
	"encoding/json"
	"fmt"
)

func ConvertDataToJSON(input interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal data to json: %v", err)
	}

	return data, nil
}
