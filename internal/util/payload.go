package util

import (
	"encoding/json"
	"io"
)

func ReadAndParseBody(body io.ReadCloser, target interface{}) error {
	defer body.Close()
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, target)
}
