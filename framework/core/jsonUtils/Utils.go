package jsonUtils

import (
	"encoding/json"
	"io"
)

func ToJson(i interface{}, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(i)
}

func FromJson(i interface{}, reader io.Reader) error{
	return json.NewDecoder(reader).Decode(i)
}
