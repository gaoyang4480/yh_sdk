package convert

import "encoding/json"

func encodeData(data map[string]string) (string, error) {
	if data == nil {
		return "", nil
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
