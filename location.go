package goeverynet

import "encoding/json"

type LocationFrame struct {
	Type   string         `json:"type"`
	Params LocationParams `json:"params"`
	Meta   Meta           `json:"meta"`
}

type LocationParams struct {
	Solutions []Solutions `json:"solutions"`
}

type Solutions struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Quality   int     `json:"quality"`
	Precision int     `json:"precision"`
	Method    string  `json:"method"`
}

// NewLocationFrame creates a new LocationFrame
func NewLocationFrame(payload []byte) (*LocationFrame, error) {
	var frame LocationFrame

	err := json.Unmarshal(payload, &frame)
	if err != nil {
		return nil, err
	}

	err = frame.Validate()
	if err != nil {
		return nil, err
	}

	return &frame, nil
}

// MarhsalJSON returns the JSON encoding of a LocationFrame
func (frame *LocationFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(frame)
}

// Validate validates a LocationFrame
func (frame *LocationFrame) Validate() error {
	return nil
}
