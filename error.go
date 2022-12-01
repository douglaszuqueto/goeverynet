package goeverynet

import "encoding/json"

type ErrorFrame struct {
	Type   string      `json:"type"`
	Params ErrorParams `json:"params"`
	Meta   ErrorMeta   `json:"meta"`
}

type ErrorParams struct {
	Message string `json:"message"`
}

type ErrorMeta struct {
	Time float64 `json:"time"`
}

// NewErrorFrame creates a new ErrorFrame from a byte array
func NewErrorFrame(data []byte) (*ErrorFrame, error) {
	var frame ErrorFrame

	err := json.Unmarshal(data, &frame)
	if err != nil {
		return nil, err
	}

	err = frame.Validate()
	if err != nil {
		return nil, err
	}

	return &frame, nil
}

// MarshalJSON returns the JSON encoding of a ErrorFrame
func (frame *ErrorFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(frame)
}

// Validate validates a ErrorFrame
func (frame *ErrorFrame) Validate() error {
	return nil
}
