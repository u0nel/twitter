package twitter

import (
	"encoding/json"
	"errors"
)

func JsonGet(data []byte, keys ...string) ([]byte, error) {
	for _, key := range keys {
		var d map[string]json.RawMessage
		err := json.Unmarshal(data, &d)
		if err != nil {
			return data, err
		}
		tdata, ok := d[key]
		if !ok {
			return data, errors.New("Key not in Json Document")
		}
		data = tdata

	}
	return data, nil
}

func JsonGetI(data []byte, i int) ([]byte, error) {
	var d []json.RawMessage
	err := json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	if len(d) <= i {
		return nil, errors.New("Index not Available")
	}
	return d[i], nil
}
