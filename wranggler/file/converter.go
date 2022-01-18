package file

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func ConvertDataToJSON(input interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return []byte{}, err
	}

	return data, err
}

func ConvertDataToCSV(inputp interface{}) ([]byte, error) {
	input, ok := inputp.([]interface{})
	if !ok {
		return []byte{}, errors.New("failed to infer data type to array of interfaces")
	}

	w := &bytes.Buffer{}
	writer := csv.NewWriter(w)
	// Because csv package does not have something like
	// json.Marshal, we'll gonna do what Thanos did.
	//
	// "Fine. I'll do it myself."

	// Create the CSV headers first
	structType := reflect.TypeOf(input[0])
	headers := make([]string, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		headers = append(headers, structType.Field(i).Tag.Get("csv"))
	}

	err := writer.Write(headers)
	if err != nil {
		return []byte{}, err
	}

	for _, inputItem := range input {
		// Struct are always in-order, so it's easy to
		// put it into the temporary
		structValue := reflect.ValueOf(inputItem)
		data := make([]string, structValue.NumField())

		for k := 0; k < structValue.NumField(); k++ {
			currentValue := structValue.Field(k)

			switch currentValue.Interface().(type) {
			case bool:
				data = append(data, strconv.FormatBool(currentValue.Bool()))
				continue
			case string:
				data = append(data, currentValue.String())
				continue
			case uint:
				data = append(data, strconv.FormatUint(currentValue.Uint(), 10))
			case int64:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case int:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case time.Time:
				t, ok := currentValue.Interface().(time.Time)
				if !ok {
					return []byte{}, fmt.Errorf("struct name of %s has a type of time.Time yet cannot be parsed", currentValue.Type().Name())
				}
				data = append(data, t.Format(time.RFC3339Nano))
				continue
			default:
				return []byte{}, fmt.Errorf("struct name of %s has a weird and unsupported type", currentValue.Type().Name())
			}
		}

		err := writer.Write(data)
		if err != nil {
			return []byte{}, err
		}
	}

	writer.Flush()
	if writer.Error() != nil {
		return []byte{}, fmt.Errorf("last csv write error: %v", err)
	}

	return w.Bytes(), nil
}
