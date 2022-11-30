package goeverynet

import "encoding/json"

func GetFrameType(frame []byte) (string, error) {
	var frameType struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(frame, &frameType)
	if err != nil {
		return "", err
	}

	return frameType.Type, nil
}
